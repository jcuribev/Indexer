package file

import (
	"archive/tar"
	"fmt"
	"os"
	"path/filepath"
)

const malformedFilesDir = "./MalformedEmails/"
const indexEmailsDir = "./IndexEmails/emails"

func WriteEmailToFile(json []byte, jsonFile *os.File) {
	jsonFile.Write(json)
	jsonFile.Write([]byte(",\n"))
}

func CreateJsonFile(fileNumber string) *os.File {

	fileDir, _ := filepath.Abs(indexEmailsDir + fileNumber + ".ndjson")
	file, err := os.Create(fileDir)

	if err != nil {
		fmt.Println(err)
	}

	return file
}

func StoreMalformedFile(malformedFile *tar.Header, fileContent []byte) {
	fileDir, _ := filepath.Abs(malformedFilesDir + malformedFile.FileInfo().Name())
	file, err := os.Create(fileDir)

	if err != nil {
		fmt.Println(err)
	} else {
		file.Write(fileContent)
	}

	file.Close()
}
