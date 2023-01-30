package file

import (
	"Indexer/email"
	"Indexer/json_manager"
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

const maxItemsPerJson = 50000

func OpenTgzFile(sourceFileDir string) {

	file, err := os.Open(sourceFileDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	gzf, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tarReader := tar.NewReader(gzf)
	readFiles(*tarReader)
}

func readFiles(tarReader tar.Reader) {

	items := 0
	fileNumber := 0
	isFirstFile := true

	jsonFile := CreateJsonFile(fileNumber)
	json_manager.InitFile(jsonFile)

	for {
		file, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		switch file.Typeflag {

		case tar.TypeDir:
			continue

		case tar.TypeReg:
			fileContent, err := io.ReadAll(&tarReader)

			if err != nil {
				fmt.Println(err)
				continue
			}

			if err := convertContent(fileContent, &isFirstFile, jsonFile); err != nil {
				StoreMalformedFile(file, fileContent)
				continue
			}

			items++

			if items >= maxItemsPerJson {
				jsonFile = nextFile(&items, &fileNumber, jsonFile, &isFirstFile)
			}

		default:
			fmt.Printf("%s : %s\n", "Cannot read this file: ", file.Name)
		}
	}

	json_manager.FinishFile(jsonFile)
	jsonFile.Close()
}

func nextFile(items *int, fileNumber *int, jsonFile *os.File, isFirstFile *bool) *os.File {

	*items = 0
	json_manager.FinishFile(jsonFile)
	jsonFile.Close()

	*fileNumber++

	fmt.Println(*fileNumber)

	jsonFile = CreateJsonFile(*fileNumber)
	json_manager.InitFile(jsonFile)

	*isFirstFile = true

	return jsonFile
}

func convertContent(fileContent []byte, isFirstFile *bool, jsonFile *os.File) error {
	email, err := email.ParseContent(string(fileContent))

	if err != nil {
		fmt.Println(err)
		return err
	}

	json := json_manager.EmailToJson(email)

	WriteEmailToFile(json, isFirstFile, jsonFile)

	return nil
}

func ReadFilesFromDirectory(dir string) []fs.FileInfo {
	fileNames, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
	}

	return fileNames
}
