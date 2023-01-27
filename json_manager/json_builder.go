package json_manager

import (
	"Indexer/email"
	"encoding/json"
	"fmt"
	"os"
)

func InitFile(jsonFile *os.File) {
	jsonFile.WriteString(`{ "index" : "emails", "records": [` + "\n")
}

func FinishFile(jsonFile *os.File) {
	jsonFile.WriteString(`]}`)
}

func EmailToJson(email email.Email) []byte {
	jsondata, err := json.Marshal(email)

	if err != nil {
		fmt.Println(err)
	}

	return jsondata
}
