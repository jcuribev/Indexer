package file

import (
	"archive/tar"
	"os"
)

func WriteEmailToFile(json []byte, jsonFile *os.File) {

	//jsonFile.WriteString(`{ "index" : "emails", "records": [` + "\n")
	//jsonFile.Write([]byte("\n"))
	jsonFile.Write(json)
	jsonFile.Write([]byte(",\n"))
}

func InitFile(jsonFile *os.File) {

}

func FinishFile(jsonFile *os.File) {

}

func StoreMalformedFile(*tar.Header) {

}
