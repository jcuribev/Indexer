package json_manager

import (
	"Indexer/email"
	"encoding/json"
	"os"
)

func InitFile(jsonFile *os.File) error {
	_, err := jsonFile.WriteString(`{ "index" : "emails", "records": [` + "\n")

	if err != nil {
		return err
	}

	return nil
}

func FinishFile(jsonFile *os.File) error {
	if _, err := jsonFile.WriteString(`]}`); err != nil {
		return err
	}

	return nil
}

func EmailToJson(email email.Email) ([]byte, error) {
	jsondata, err := json.Marshal(email)

	if err != nil {
		return nil, err
	}

	return jsondata, err
}
