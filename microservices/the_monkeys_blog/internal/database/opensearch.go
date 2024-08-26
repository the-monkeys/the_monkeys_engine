package database

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sirupsen/logrus"
	"github.com/the-monkeys/the_monkeys/apis/serviceconn/gateway_blog/pb"
	"github.com/the-monkeys/the_monkeys/microservices/the_monkeys_blog/internal/constants"
)

type ElasticsearchStorage interface {
	DraftABlog(ctx context.Context, blog *pb.DraftBlogRequest) (*esapi.Response, error)
	GetDraftBlogsByOwnerAccountID(ctx context.Context, ownerAccountID string) (*pb.GetDraftBlogsRes, error)
	DoesBlogExist(ctx context.Context, blogID string) (bool, error)
	PublishBlogById(ctx context.Context, blogId string) (*esapi.Response, error)
	GetPublishedBlogByTagsName(ctx context.Context, tags ...string) (*pb.GetBlogsByTagsNameRes, error)
	GetPublishedBlogById(ctx context.Context, id string) (*pb.GetBlogByIdRes, error)
	AchieveAPublishedBlogById(ctx context.Context, blogId string) (*esapi.Response, error)
	DeleteABlogById(ctx context.Context, blogId string) (*esapi.Response, error)
	GetLast100BlogsLatestFirst(ctx context.Context) (*pb.GetBlogsByTagsNameRes, error)
}

type elasticsearchStorage struct {
	client *elasticsearch.Client
	log    *logrus.Logger
}

func NewElasticsearchClient(url, username, password string, log *logrus.Logger) (ElasticsearchStorage, error) {
	client, err := NewESClient(url, username, password)
	if err != nil {
		log.Errorf("Failed to connect to Elasticsearch instance, error: %+v", err)
		return nil, err
	}

	return &elasticsearchStorage{
		client: client,
		log:    log,
	}, nil
}

func (es *elasticsearchStorage) DraftABlog(ctx context.Context, blog *pb.DraftBlogRequest) (*esapi.Response, error) {
	bs, err := json.Marshal(blog)
	if err != nil {
		es.log.Errorf("DraftABlog: cannot marshal the blog, error: %v", err)
		return nil, err
	}

	document := strings.NewReader(string(bs))

	req := esapi.IndexRequest{
		Index:      constants.ElasticsearchBlogIndex,
		DocumentID: blog.BlogId,
		Body:       document,
	}

	insertResponse, err := req.Do(ctx, es.client)
	if err != nil {
		es.log.Errorf("DraftABlog: error while indexing blog, error: %+v", err)
		return insertResponse, err
	}

	if insertResponse.IsError() {
		err = fmt.Errorf("DraftABlog: error indexing blog, response: %+v", insertResponse)
		es.log.Error(err)
		return insertResponse, err
	}

	es.log.Infof("DraftABlog: successfully created blog for user: %s, response: %+v", blog.OwnerAccountId, insertResponse)
	return insertResponse, nil
}

func (es *elasticsearchStorage) GetDraftBlogsByOwnerAccountID(ctx context.Context, ownerAccountID string) (*pb.GetDraftBlogsRes, error) {
	// Ensure ownerAccountID is properly set
	if ownerAccountID == "" {
		es.log.Error("GetDraftBlogsByOwnerAccountID: ownerAccountID is empty")
		return nil, fmt.Errorf("ownerAccountID cannot be empty")
	}

	// Build the query to search for draft blogs by owner_account_id
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"owner_account_id.keyword": ownerAccountID,
						},
					},
					{
						"term": map[string]interface{}{
							"is_draft": true,
						},
					},
				},
			},
		},
	}

	// Marshal the query to JSON
	bs, err := json.Marshal(query)
	if err != nil {
		es.log.Errorf("GetDraftBlogsByOwnerAccountID: cannot marshal the query, error: %v", err)
		return nil, err
	}

	// Print the query for debugging
	es.log.Infof("Executing query: %s", string(bs))

	// Create a new search request with the query
	req := esapi.SearchRequest{
		Index: []string{constants.ElasticsearchBlogIndex},
		Body:  strings.NewReader(string(bs)),
	}

	// Execute the search request
	res, err := req.Do(ctx, es.client)
	if err != nil {
		es.log.Errorf("GetDraftBlogsByOwnerAccountID: error executing search request, error: %+v", err)
		return nil, err
	}
	defer res.Body.Close()

	// Check if the response indicates an error
	if res.IsError() {
		err = fmt.Errorf("GetDraftBlogsByOwnerAccountID: search query failed, response: %+v", res)
		es.log.Error(err)
		return nil, err
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		es.log.Errorf("GetDraftBlogsByOwnerAccountID: error reading response body, error: %v", err)
		return nil, err
	}

	// Print the response body for debugging
	// es.log.Infof("Search response body: %s", string(bodyBytes))

	// Parse the response body
	var esResponse map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &esResponse); err != nil {
		es.log.Errorf("GetDraftBlogsByOwnerAccountID: error decoding response body, error: %v", err)
		return nil, err
	}

	// Extract the hits from the response
	hits, ok := esResponse["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		err := fmt.Errorf("GetDraftBlogsByOwnerAccountID: failed to parse hits from response")
		es.log.Error(err)
		return nil, err
	}

	// Convert the hits to a slice of DraftBlogRequest
	var blogs = &pb.GetDraftBlogsRes{
		Blogs: make([]*pb.GetBlogs, 0, len(hits)),
	}
	for _, hit := range hits {
		hitSource := hit.(map[string]interface{})["_source"]
		hitBytes, err := json.Marshal(hitSource)
		if err != nil {
			es.log.Errorf("GetDraftBlogsByOwnerAccountID: error marshaling hit source, error: %v", err)
			continue
		}

		var blog pb.GetBlogs
		if err := json.Unmarshal(hitBytes, &blog); err != nil {
			es.log.Errorf("GetDraftBlogsByOwnerAccountID: error unmarshaling hit to DraftBlogRequest, error: %v", err)
			continue
		}
		blogs.Blogs = append(blogs.Blogs, &blog)
	}

	es.log.Infof("GetDraftBlogsByOwnerAccountID: successfully fetched %d draft blogs for owner_account_id: %s", len(blogs.Blogs), ownerAccountID)
	return blogs, nil
}

func (es *elasticsearchStorage) DoesBlogExist(ctx context.Context, blogID string) (bool, error) {
	// Ensure blogID is not empty
	if blogID == "" {
		es.log.Error("DoesBlogExist: blogID is empty")
		return false, fmt.Errorf("blogID cannot be empty")
	}

	// Create a Get request to check if the document exists
	req := esapi.GetRequest{
		Index:      constants.ElasticsearchBlogIndex,
		DocumentID: blogID,
	}

	// Execute the Get request
	getResponse, err := req.Do(ctx, es.client)
	if err != nil {
		es.log.Errorf("DoesBlogExist: error executing Get request, error: %+v", err)
		return false, err
	}
	defer getResponse.Body.Close()

	// Check if the response indicates the document exists
	if getResponse.StatusCode == http.StatusOK {
		es.log.Infof("DoesBlogExist: blog with id %s exists", blogID)
		return true, nil
	} else if getResponse.StatusCode == http.StatusNotFound {
		es.log.Infof("DoesBlogExist: blog with id %s does not exist", blogID)
		return false, nil
	}

	// If the response is something else, log it as an error
	err = fmt.Errorf("DoesBlogExist: unexpected status code %d", getResponse.StatusCode)
	es.log.Error(err)
	return false, err
}

func (es *elasticsearchStorage) PublishBlogById(ctx context.Context, blogId string) (*esapi.Response, error) {
	// Ensure blogId is not empty
	if blogId == "" {
		es.log.Error("PublishBlogById: blogId is empty")
		return nil, fmt.Errorf("blogId cannot be empty")
	}

	// Build the update query to set is_draft to false
	updateScript := map[string]interface{}{
		"script": map[string]interface{}{
			"source": "ctx._source.is_draft = false",
		},
	}

	// Marshal the update script to JSON
	bs, err := json.Marshal(updateScript)
	if err != nil {
		es.log.Errorf("PublishBlogById: cannot marshal the update script, error: %v", err)
		return nil, err
	}

	// Create an update request
	req := esapi.UpdateRequest{
		Index:      constants.ElasticsearchBlogIndex,
		DocumentID: blogId,
		Body:       strings.NewReader(string(bs)),
	}

	// Execute the update request
	updateResponse, err := req.Do(ctx, es.client)
	if err != nil {
		es.log.Errorf("PublishBlogById: error executing update request, error: %+v", err)
		return updateResponse, err
	}
	defer updateResponse.Body.Close()

	// Check if the response indicates an error
	if updateResponse.IsError() {
		err = fmt.Errorf("PublishBlogById: update query failed, response: %+v", updateResponse)
		es.log.Error(err)
		return updateResponse, err
	}

	es.log.Infof("PublishBlogById: successfully published blog with id: %s", blogId)
	return updateResponse, nil
}

func (es *elasticsearchStorage) GetPublishedBlogByTagsName(ctx context.Context, tags ...string) (*pb.GetBlogsByTagsNameRes, error) {
	// Ensure at least one tag is provided
	if len(tags) == 0 {
		es.log.Error("GetPublishedBlogByTagsName: no tags provided")
		return nil, fmt.Errorf("at least one tag must be provided")
	}

	// Build the query to search for published blogs by tags
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"terms": map[string]interface{}{
							"tags.keyword": tags,
						},
					},
					{
						"term": map[string]interface{}{
							"is_draft": false,
						},
					},
				},
			},
		},
	}

	// Marshal the query to JSON
	bs, err := json.Marshal(query)
	if err != nil {
		es.log.Errorf("GetPublishedBlogByTagsName: cannot marshal the query, error: %v", err)
		return nil, err
	}

	// Print the query for debugging
	es.log.Infof("Executing query: %s", string(bs))

	// Create a new search request with the query
	req := esapi.SearchRequest{
		Index: []string{constants.ElasticsearchBlogIndex},
		Body:  strings.NewReader(string(bs)),
	}

	// Execute the search request
	res, err := req.Do(ctx, es.client)
	if err != nil {
		es.log.Errorf("GetPublishedBlogByTagsName: error executing search request, error: %+v", err)
		return nil, err
	}
	defer res.Body.Close()

	// Check if the response indicates an error
	if res.IsError() {
		err = fmt.Errorf("GetPublishedBlogByTagsName: search query failed, response: %+v", res)
		es.log.Error(err)
		return nil, err
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		es.log.Errorf("GetPublishedBlogByTagsName: error reading response body, error: %v", err)
		return nil, err
	}

	// Parse the response body
	var esResponse map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &esResponse); err != nil {
		es.log.Errorf("GetPublishedBlogByTagsName: error decoding response body, error: %v", err)
		return nil, err
	}

	// Extract the hits from the response
	hits, ok := esResponse["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		err := fmt.Errorf("GetPublishedBlogByTagsName: failed to parse hits from response")
		es.log.Error(err)
		return nil, err
	}

	// Convert the hits to a slice of GetBlogs
	var blogs = &pb.GetBlogsByTagsNameRes{
		TheBlogs: make([]*pb.GetBlogsByTags, 0, len(hits)),
	}
	for _, hit := range hits {
		hitSource := hit.(map[string]interface{})["_source"]
		hitBytes, err := json.Marshal(hitSource)
		if err != nil {
			es.log.Errorf("GetPublishedBlogByTagsName: error marshaling hit source, error: %v", err)
			continue
		}

		var blog pb.GetBlogsByTags
		if err := json.Unmarshal(hitBytes, &blog); err != nil {
			es.log.Errorf("GetPublishedBlogByTagsName: error unmarshaling hit to GetBlogs, error: %v", err)
			continue
		}
		blogs.TheBlogs = append(blogs.TheBlogs, &blog)
	}

	es.log.Infof("GetPublishedBlogByTagsName: successfully fetched %d published blogs for tags: %v", len(blogs.TheBlogs), tags)
	return blogs, nil
}

func (es *elasticsearchStorage) GetPublishedBlogById(ctx context.Context, id string) (*pb.GetBlogByIdRes, error) {
	// Ensure id is not empty
	if id == "" {
		es.log.Error("GetPublishedBlogById: id is empty")
		return nil, fmt.Errorf("blog id cannot be empty")
	}

	// Build the query to search for a published blog by id
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"blog_id.keyword": id,
						},
					},
					{
						"term": map[string]interface{}{
							"is_draft": false,
						},
					},
				},
			},
		},
	}

	// Marshal the query to JSON
	bs, err := json.Marshal(query)
	if err != nil {
		es.log.Errorf("GetPublishedBlogById: cannot marshal the query, error: %v", err)
		return nil, err
	}

	// Print the query for debugging
	es.log.Infof("Executing query: %s", string(bs))

	// Create a new search request with the query
	req := esapi.SearchRequest{
		Index: []string{constants.ElasticsearchBlogIndex},
		Body:  strings.NewReader(string(bs)),
	}

	// Execute the search request
	res, err := req.Do(ctx, es.client)
	if err != nil {
		es.log.Errorf("GetPublishedBlogById: error executing search request, error: %+v", err)
		return nil, err
	}
	defer res.Body.Close()

	// Check if the response indicates an error
	if res.IsError() {
		err = fmt.Errorf("GetPublishedBlogById: search query failed, response: %+v", res)
		es.log.Error(err)
		return nil, err
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		es.log.Errorf("GetPublishedBlogById: error reading response body, error: %v", err)
		return nil, err
	}

	// Parse the response body
	var esResponse map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &esResponse); err != nil {
		es.log.Errorf("GetPublishedBlogById: error decoding response body, error: %v", err)
		return nil, err
	}

	// Extract the hits from the response
	hits, ok := esResponse["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok || len(hits) == 0 {
		es.log.Infof("GetPublishedBlogById: no published blog found with id: %s", id)
		return nil, nil
	}

	// Convert the first hit to GetBlogByIdRes
	hitSource := hits[0].(map[string]interface{})["_source"]
	hitBytes, err := json.Marshal(hitSource)
	if err != nil {
		es.log.Errorf("GetPublishedBlogById: error marshaling hit source, error: %v", err)
		return nil, err
	}

	var blog pb.GetBlogByIdRes
	if err := json.Unmarshal(hitBytes, &blog); err != nil {
		es.log.Errorf("GetPublishedBlogById: error unmarshaling hit to GetBlogByIdRes, error: %v", err)
		return nil, err
	}

	es.log.Infof("GetPublishedBlogById: successfully fetched published blog with id: %s", id)
	return &blog, nil
}

// AchieveAPublishedBlogById archives a published blog by setting an "is_archived" field to true
func (es *elasticsearchStorage) AchieveAPublishedBlogById(ctx context.Context, blogId string) (*esapi.Response, error) {
	// Ensure blogId is not empty
	if blogId == "" {
		es.log.Error("AchieveAPublishedBlogById: blogId is empty")
		return nil, fmt.Errorf("blogId cannot be empty")
	}

	// Build the update query to set is_archived to true
	updateScript := map[string]interface{}{
		"script": map[string]interface{}{
			"source": "ctx._source.is_archived = true",
		},
	}

	// Marshal the update script to JSON
	bs, err := json.Marshal(updateScript)
	if err != nil {
		es.log.Errorf("AchieveAPublishedBlogById: cannot marshal the update script, error: %v", err)
		return nil, err
	}

	// Create an update request
	req := esapi.UpdateRequest{
		Index:      constants.ElasticsearchBlogIndex,
		DocumentID: blogId,
		Body:       strings.NewReader(string(bs)),
	}

	// Execute the update request
	updateResponse, err := req.Do(ctx, es.client)
	if err != nil {
		es.log.Errorf("AchieveAPublishedBlogById: error executing update request, error: %+v", err)
		return updateResponse, err
	}
	defer updateResponse.Body.Close()

	// Check if the response indicates an error
	if updateResponse.IsError() {
		err = fmt.Errorf("AchieveAPublishedBlogById: update query failed, response: %+v", updateResponse)
		es.log.Error(err)
		return updateResponse, err
	}

	es.log.Infof("AchieveAPublishedBlogById: successfully archived blog with id: %s", blogId)
	return updateResponse, nil
}

// DeleteABlogById deletes a blog by its ID
func (es *elasticsearchStorage) DeleteABlogById(ctx context.Context, blogId string) (*esapi.Response, error) {
	// Ensure blogId is not empty
	if blogId == "" {
		es.log.Error("DeleteABlogById: blogId is empty")
		return nil, fmt.Errorf("blogId cannot be empty")
	}

	// Create a Delete request
	req := esapi.DeleteRequest{
		Index:      constants.ElasticsearchBlogIndex,
		DocumentID: blogId,
	}

	// Execute the delete request
	deleteResponse, err := req.Do(ctx, es.client)
	if err != nil {
		es.log.Errorf("DeleteABlogById: error executing delete request, error: %+v", err)
		return deleteResponse, err
	}
	defer deleteResponse.Body.Close()

	// Check if the response indicates an error
	if deleteResponse.IsError() {
		err = fmt.Errorf("DeleteABlogById: delete query failed, response: %+v", deleteResponse)
		es.log.Error(err)
		return deleteResponse, err
	}

	es.log.Infof("DeleteABlogById: successfully deleted blog with id: %s", blogId)
	return deleteResponse, nil
}

// GetLast100BlogsLatestFirst retrieves the last 100 blogs sorted by the latest first
func (es *elasticsearchStorage) GetLast100BlogsLatestFirst(ctx context.Context) (*pb.GetBlogsByTagsNameRes, error) {
	// Build the query to retrieve the last 100 blogs, sorted by a timestamp field in descending order
	query := map[string]interface{}{
		"sort": []map[string]interface{}{
			{
				"timestamp": map[string]string{
					"order": "desc",
				},
			},
		},
		"size": 100,
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"is_draft": false,
			},
		},
	}

	// Marshal the query to JSON
	bs, err := json.Marshal(query)
	if err != nil {
		es.log.Errorf("GetLast100BlogsLatestFirst: cannot marshal the query, error: %v", err)
		return nil, err
	}

	// Create a new search request with the query
	req := esapi.SearchRequest{
		Index: []string{constants.ElasticsearchBlogIndex},
		Body:  strings.NewReader(string(bs)),
	}

	// Execute the search request
	res, err := req.Do(ctx, es.client)
	if err != nil {
		es.log.Errorf("GetLast100BlogsLatestFirst: error executing search request, error: %+v", err)
		return nil, err
	}
	defer res.Body.Close()

	// Check if the response indicates an error
	if res.IsError() {
		err = fmt.Errorf("GetLast100BlogsLatestFirst: search query failed, response: %+v", res)
		es.log.Error(err)
		return nil, err
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		es.log.Errorf("GetLast100BlogsLatestFirst: error reading response body, error: %v", err)
		return nil, err
	}

	// Parse the response body
	var esResponse map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &esResponse); err != nil {
		es.log.Errorf("GetLast100BlogsLatestFirst: error decoding response body, error: %v", err)
		return nil, err
	}

	// Extract the hits from the response
	hits, ok := esResponse["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		err := fmt.Errorf("GetLast100BlogsLatestFirst: failed to parse hits from response")
		es.log.Error(err)
		return nil, err
	}

	// Convert the hits to a slice of GetBlogs
	var blogs = &pb.GetBlogsByTagsNameRes{
		TheBlogs: make([]*pb.GetBlogsByTags, 0, len(hits)),
	}
	for _, hit := range hits {
		hitSource := hit.(map[string]interface{})["_source"]
		hitBytes, err := json.Marshal(hitSource)
		if err != nil {
			es.log.Errorf("GetLast100BlogsLatestFirst: error marshaling hit source, error: %v", err)
			continue
		}

		var blog pb.GetBlogsByTags
		if err := json.Unmarshal(hitBytes, &blog); err != nil {
			es.log.Errorf("GetLast100BlogsLatestFirst: error unmarshaling hit to GetBlogsByTags, error: %v", err)
			continue
		}
		blogs.TheBlogs = append(blogs.TheBlogs, &blog)
	}

	es.log.Infof("GetLast100BlogsLatestFirst: successfully fetched last 100 blogs sorted by latest first")
	return blogs, nil
}
