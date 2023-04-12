package server

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/89minutes/the_new_project/services/file_server/service/pb"
	"github.com/89minutes/the_new_project/services/file_server/service/utils"
	"github.com/sirupsen/logrus"
)

type FileService struct {
	path string
	pb.UnimplementedUploadBlogFileServer
}

func NewFileService(path string) *FileService {
	return &FileService{path: path}
}

func (fs *FileService) UploadBlogFile(stream pb.UploadBlogFile_UploadBlogFileServer) error {
	var byteSlice []byte
	var blogId string
	var fileName string
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		byteSlice = append(byteSlice, chunk.Data...)
		blogId = chunk.BlogId
		fileName = chunk.FileName
	}
	logrus.Infof("Uploading a file for blog id: %v", blogId)

	fileName = utils.RemoveSpecialChar(fileName)
	dirPath, filePath := utils.ConstructPath(fs.path, blogId, fileName)

	// Check if directory exists, if not create it
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		logrus.Infof("the directory, %s doesn't exists", dirPath)

		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			logrus.Errorf("cannot create a directory for this blog id: %s", blogId)
			return err
		}
	}

	// Check if file exists, if not create it with sample data
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		logrus.Infof("the file, %s doesn't exists", filePath)

		err := ioutil.WriteFile(filePath, byteSlice, 0644)
		if err != nil {
			logrus.Errorf("cannot create a file for this blog id: %s", blogId)
			return err
		}
	}

	logrus.Infof("done uploading file: %s", filePath)
	return stream.SendAndClose(&pb.UploadBlogFileRes{
		Status:      http.StatusOK,
		NewFileName: fileName,
	})
}

// TODO: return error 404 if file doesn't exist
func (fs *FileService) GetBlogFile(req *pb.GetBlogFileReq, stream pb.UploadBlogFile_GetBlogFileServer) error {
	fileName := filepath.Join(fs.path, req.BlogId, req.FileName)
	logrus.Infof("there is a request to retrieve the file, %s", fileName)

	rawFileName := strings.ReplaceAll(fileName, "\n", "")
	fileBytes, err := ioutil.ReadFile(rawFileName)
	if err != nil {
		logrus.Errorf("cannot read the file: %s, error: %v", fileName, fileBytes)
		return err
	}

	if err := stream.Send(&pb.GetBlogFileRes{
		Data: fileBytes,
	}); err != nil {
		logrus.Errorf("error while sending stream, error %+v", err)
	}

	return nil
}

func (fs *FileService) DeleteBlogFile(ctx context.Context, req *pb.DeleteBlogFileReq) (*pb.DeleteBlogFileRes, error) {
	filePath := filepath.Join(fs.path, req.BlogId, req.FileName)

	logrus.Infof("there is a request to delete the file, %s", filePath)

	if err := os.Remove(filePath); err != nil {
		return nil, err
	}

	return &pb.DeleteBlogFileRes{
		Message: "successfully deleted",
		Status:  http.StatusOK,
	}, nil
}
