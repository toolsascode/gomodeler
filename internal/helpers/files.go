package helpers

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func ReadFilesDir(path string) []string {

	var fileList []string

	if len(path) == 0 {
		return nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Debugf("Reading file: %s", file.Name())
		fileList = append(fileList, fmt.Sprintf("%s/%s", path, file.Name()))
	}

	return fileList
}
