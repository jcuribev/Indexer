package main

import (
	"Indexer/Indexer"
)

func main() {

	// flag.Parse()
	// sourceFile := flag.Arg(0)

	// if sourceFile == "" {
	// 	fmt.Println("Missing file argument.")
	// 	os.Exit(1)
	// }

	// fmt.Println("Selected tgz:", flag.Arg(0))

	// file.OpenTgzFile(sourceFile)

	Indexer.IndexEmails()
}