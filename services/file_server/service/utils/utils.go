package utils

import (
	"path/filepath"
	"regexp"

	"github.com/sirupsen/logrus"
)

func RemoveSpecialChar(val string) string {
	// Define regular expression to match all special characters
	reg, err := regexp.Compile("[^a-zA-Z0-9\\.]+")
	if err != nil {
		logrus.Error("cannot compile regexp")
	}

	// Remove special characters from the input string
	output := reg.ReplaceAllString(val, "")

	return output
}

func ConstructPath(basePath, blogId, fileName string) (string, string) {
	return filepath.Join(basePath, blogId), filepath.Join(basePath, blogId, fileName)
}
