package file

import (
	"archive/tar"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const malformedFilesDir = "./MalformedEmails/"
const indexEmailsDir = "./IndexEmails/emails"

func WriteEmailToFile(json []byte, isFirstFile *bool, jsonFile *os.File) {

	if *isFirstFile == true {
		*isFirstFile = false
	} else {
		jsonFile.Write([]byte(",\n"))
	}
	jsonFile.Write(json)

}

func CreateJsonFile(fileNumber int) *os.File {

	fileDir, _ := filepath.Abs(indexEmailsDir + strconv.Itoa(fileNumber) + ".ndjson")
	file, err := os.Create(fileDir)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}

func StoreMalformedFile(malformedFile *tar.Header, fileContent []byte) {

	fileDir, _ := filepath.Abs(malformedFilesDir + malformedFile.FileInfo().Name())
	file, err := os.Create(fileDir)

	defer file.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		file.Write(fileContent)
	}
}
