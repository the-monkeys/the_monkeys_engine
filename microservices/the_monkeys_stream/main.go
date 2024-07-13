package main

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/stream/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filePath := filepath.Join("path_to_your_videos", filename)

		c.Header("Content-Type", "video/mp4")
		c.Header("Transfer-Encoding", "chunked")

		c.Stream(func(w io.Writer) bool {
			file, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			_, err = io.Copy(w, file)
			if err == nil {
				return true
			}

			return true
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("cannot start websocket at :8080: %v", err)
	}

}
