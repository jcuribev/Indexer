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
	"strconv"
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
	jsonFile := CreateJsonFile(strconv.Itoa(fileNumber))
	json_manager.InitFile(jsonFile)

	for {
		file, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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

			if err := convertContent(string(fileContent), jsonFile); err != nil {
				StoreMalformedFile(file, fileContent)
				continue
			}

			items++
			if items >= maxItemsPerJson {
				items = 0
				fileNumber++
				json_manager.FinishFile(jsonFile)

				jsonFile.Close()

				jsonFile = CreateJsonFile(strconv.Itoa(fileNumber))

				json_manager.InitFile(jsonFile)
			}

		default:
			fmt.Printf("%s : %s\n", "Cannot read this file: ", file.Name)
		}
	}

	json_manager.FinishFile(jsonFile)
}

func convertContent(fileContent string, jsonFile *os.File) error {
	email, err := email.ParseContent(string(fileContent))

	if err != nil {
		fmt.Println(err)
		return err
	}

	json := json_manager.EmailToJson(email)

	WriteEmailToFile(json, jsonFile)

	return nil
}

func ReadFilesFromDirectory(dir string) []fs.FileInfo {
	fileNames, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Println(err)
	}

	return fileNames
}
