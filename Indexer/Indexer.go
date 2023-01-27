package Indexer

import (
	"Indexer/file"
	"net/http"
	"os"
)

const bulkAddress = "http://localhost:4080/api/_bulkv2"
const directory = "./IndexEmails/"

func IndexEmails() {

	filesInfo := file.ReadFilesFromDirectory(directory)

	for i := range filesInfo {

		file, err := os.Open(directory + filesInfo[i].Name())
		if err != nil {
			println(err)
		}
		defer file.Close()

		request, err := http.NewRequest("POST", bulkAddress, file)

		if err != nil {
			println(err)
		}

		request.SetBasicAuth("admin", "Complexpass#123")

		response, err := http.DefaultClient.Do(request)

		if err != nil {
			println(err)
		}
		defer response.Body.Close()
	}
}
