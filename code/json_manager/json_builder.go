package json_manager

import (
	"Indexer/email"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func CreateJsonFile(fileNumber string) *os.File {

	fileDir, _ := filepath.Abs("../IndexEmails/emails" + fileNumber + ".ndjson")

	fmt.Printf(fileDir, "\n")

	file, err := os.Create(fileDir)

	if err != nil {
		fmt.Println(err)
	}

	return file
}

func EmailToJson(email email.Email) []byte {
	jsondata, err := json.Marshal(email)

	if err != nil {
		fmt.Println(err)
	}

	return jsondata
}
