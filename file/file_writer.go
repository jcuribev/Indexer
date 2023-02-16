package file

import (
	"Indexer/json_manager"
	"archive/tar"
	"os"
	"path/filepath"
	"strconv"
)

const malformedFilesDir = "./MalformedEmails/"
const indexEmailsDir = "./IndexEmails/emails"

func WriteEmailToFile(json []byte, isFirstFile *bool, jsonFile *os.File) error {

	if *isFirstFile == true {
		*isFirstFile = false
		return nil
	}

	if _, err := jsonFile.Write([]byte(",\n")); err != nil {
		return err
	}

	if _, err := jsonFile.Write(json); err != nil {
		return err
	}

	return nil
}

func CreateJsonFile(fileNumber int) (*os.File, error) {

	fileDir, err := filepath.Abs(indexEmailsDir + strconv.Itoa(fileNumber) + ".ndjson")

	if err != nil {
		return nil, err
	}

	jsonFile, err := os.Create(fileDir)

	if err != nil {
		return nil, err
	}

	if err := json_manager.InitFile(jsonFile); err != nil {
		return nil, err
	}

	return jsonFile, nil
}

func StoreMalformedFile(tarHeader *tar.Header, fileContent []byte) error {

	fileDir, err := filepath.Abs(malformedFilesDir + tarHeader.Name)

	if err != nil {
		return err
	}

	file, err := os.Create(fileDir)

	if err != nil {
		return err
	}

	_, err = file.Write(fileContent)

	if err != nil {
		return err
	}

	file.Close()

	return nil
}

func CreateProfileFile(fileName string) (*os.File, error) {
	f, err := os.Create(fileName)

	if err != nil {
		return nil, err
	}
	return f, nil
}
