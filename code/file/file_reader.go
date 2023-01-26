package file

import (
	"Indexer/email"
	"Indexer/json_manager"
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strconv"
)

const MaxItemsPerJson = 50000

func OpenTgzFile(sourceFile string) {

	file, err := os.Open(sourceFile)
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
	jsonFile := json_manager.CreateJsonFile(strconv.Itoa(fileNumber))
	InitFile(jsonFile)

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
				StoreMalformedFile(file)
				continue
			}

			items++
			if items >= MaxItemsPerJson {
				items = 0
				fileNumber++
				FinishFile(jsonFile)

				jsonFile = json_manager.CreateJsonFile(strconv.Itoa(fileNumber))
			}

		default:
			fmt.Printf("%s : %s\n", "Cannot read this file: ", file.Name)
		}
	}
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
