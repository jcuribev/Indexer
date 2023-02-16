package file

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

func GetTgzReader(sourceFileDir string) (*tar.Reader, error) {

	file, err := os.Open(sourceFileDir)
	if err != nil {
		return nil, err
	}

	gzf, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}

	tarReader := tar.NewReader(gzf)
	return tarReader, nil
}

func ReadFileContent(tarReader *tar.Reader) ([]byte, error) {

	fileContent, err := io.ReadAll(tarReader)

	if err != nil {
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
