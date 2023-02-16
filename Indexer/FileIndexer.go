package Indexer

import (
	"Indexer/email"
	"Indexer/file"
	"Indexer/json_manager"
	"archive/tar"
	"io"
	"os"
)

type FileInformation struct {
	entriesWritten int
	fileNumber     int
	isFirstEntry   bool
	jsonFile       *os.File
}

const maxEntriesPerJson = 50000

func IterateTarReader(tarReader *tar.Reader) error {

	fileInfo := FileInformation{
		entriesWritten: 0,
		fileNumber:     0,
		isFirstEntry:   true,
		jsonFile:       nil,
	}

	for {
		tarHeader, err := tarReader.Next()

		if err == io.EOF {
			EndJsonFile(fileInfo.jsonFile)
			break
		}

		if err != nil {
			return err
		}

		if fileInfo.jsonFile == nil {
			fileInfo.jsonFile, err = file.CreateJsonFile(fileInfo.fileNumber)

			if err != nil {
				return err
			}
		}

		if tarHeader.Typeflag != tar.TypeReg {
			//fmt.Printf("err: %v\n", errors.New("Unsupported file"))
			continue
		}

		if err := HandleFile(tarReader, &fileInfo); err == nil {
			fileInfo.isFirstEntry = false
			fileInfo.entriesWritten++
		} else if err := HandleBadFile(tarReader, tarHeader); err != nil {
			return err
		}

		if fileInfo.entriesWritten != maxEntriesPerJson {
			continue
		}

		if err := EndJsonFile(fileInfo.jsonFile); err != nil {
			return err
		}
		NewFileInformation(&fileInfo)
	}

	return nil
}

func EndJsonFile(file *os.File) error {

	defer file.Close()

	if err := json_manager.FinishFile(file); err != nil {
		return err
	}

	return nil
}

func NewFileInformation(fileInfo *FileInformation) {

	fileInfo.entriesWritten = 0
	fileInfo.fileNumber = fileInfo.fileNumber + 1
	fileInfo.isFirstEntry = true
	fileInfo.jsonFile = nil
}

func HandleFile(tarReader *tar.Reader, fileInfo *FileInformation) error {
	fileContent, err := file.ReadFileContent(tarReader)

	if err != nil {
		return err
	}

	email, err := email.FileContentToEmail(string(fileContent))

	if err != nil {
		return err
	}

	jsonEmail, err := json_manager.EmailToJson(email)

	if err != nil {
		return err
	}

	if err := file.WriteEmailToFile(jsonEmail, &fileInfo.isFirstEntry, fileInfo.jsonFile); err != nil {
		return err
	}

	return nil
}

func HandleBadFile(tarReader *tar.Reader, tarHeader *tar.Header) error {
	fileContent, err := file.ReadFileContent(tarReader)

	if err != nil {
		return err
	}

	if err := file.StoreMalformedFile(tarHeader, fileContent); err != nil {
		return err
	}

	return nil
}
