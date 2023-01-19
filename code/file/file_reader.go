package file

import (
	"Indexer/email"
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

const numberOfHeaders = 15

func ProcessFile(sourceFile string) {

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

	defer gzf.Close()

	tarReader := tar.NewReader(gzf)

	i := 0

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
			fileContent, err := io.ReadAll(tarReader)
			if err != nil {
				fmt.Println(err)
				os.Exit((1))
			}

			email.ParseContent(string(fileContent))

			if i == 1 {
				os.Exit(0)
			}

		default:
			fmt.Printf("%s : %s\n", "Cannot read this file: ", file.Name)
		}

		i++
	}
}
