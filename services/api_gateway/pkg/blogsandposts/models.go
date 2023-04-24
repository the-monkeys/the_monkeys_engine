package blogsandposts

type CreatePostRequestBody struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Author    string   `json:"author"`
	AuthorId  string   `json:"author_id"`
	Published bool     `json:"published"`
	Tags      []string `json:"tags"`
}

type Post struct {
	Id          string   `json:"id"`
	HTMLContent string   `json:"html_content"`
	RawContent  string   `json:"raw_content"`
	AuthorName  string   `json:"author_name"`
	AuthorEmail string   `json:"author_email"`
	Tags        []string `json:"tags"`
}

type EditArticleRequestBody struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type Tag struct {
	TagName string `json:"tag_name"`
}
