package service

// func ParseToStruct(result models.Last100Articles) []pb.GetBlogsResponse {
// 	var resp []pb.GetBlogsResponse

// 	for _, val := range result.Hits.Hits {

// 		res := pb.GetBlogsResponse{
// 			Id:     val.Source.ID,
// 			Title:  val.Source.Title,
// 			Author: val.Source.Author,
// 			// AuthorEmail: val.Source.AuthorEmail,
// 			CreateTime: timestamppb.New(val.Source.CreateTime),
// 			QuickRead:  val.Source.QuickRead,
// 		}
// 		resp = append(resp, res)
// 	}

// 	return resp
// }

// func partialOrAllUpdate(method string, existingArt *pb.GetBlogByIdResp, reqArt *pb.EditBlogReq) *pb.EditBlogReq {
// 	procdArt := &pb.EditBlogReq{Id: reqArt.Id}

// 	if method == http.MethodPatch {
// 		if reqArt.Title == "" {
// 			procdArt.Title = existingArt.Title
// 		} else {
// 			procdArt.Title = reqArt.Title
// 		}
// 		if reqArt.Content == "" {
// 			procdArt.Content = existingArt.Content
// 		} else {
// 			procdArt.Content = reqArt.Content
// 		}
// 		if len(reqArt.Tags) == 0 {
// 			procdArt.Tags = existingArt.Tags
// 		} else {
// 			procdArt.Tags = reqArt.Tags
// 		}
// 	} else {
// 		procdArt.Title = reqArt.Title
// 		procdArt.Content = reqArt.Content
// 		procdArt.Tags = reqArt.Tags
// 	}

// 	return procdArt
// }