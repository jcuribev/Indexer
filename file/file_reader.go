package file

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

func OpenSourceFile(sourceFile string) (*os.File, error) {
	return os.Open(sourceFile)
}

func GetTgzReader(tgzFile *os.File) (*tar.Reader, error) {

	gZipReader, err := gzip.NewReader(tgzFile)
	if err != nil {
		return nil, err
	}

	tarReader := tar.NewReader(gZipReader)

	return tarReader, nil
}

func ReadFileContent(tarReader *tar.Reader) ([]byte, error) {

	fileContent, err := io.ReadAll(tarReader)

	if err != nil {
		fmt.Println("Read file content failed")
		return nil, err
	}

	return fileContent, nil
}

func ReadFilesFromDirectory(dir string) ([]fs.FileInfo, error) {
	fileNames, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	return fileNames, nil
}
