package main

import (
	"Indexer/Indexer"
	"Indexer/file"
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {

	if err := initializeProgram(); err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
}

func initializeProgram() error {

	sourceFile := flag.String("file", "", "read file")
	cpuProfile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memProfile := flag.String("memprofile", "", "write memory profile to `file`")
	flag.Parse()
	println(*cpuProfile)

	cpuFile, err := handleProfileFlag(*cpuProfile)
	if err != nil {
		return err
	}
	defer cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		return errors.New("could not start CPU profile: " + err.Error())
	}
	defer pprof.StopCPUProfile()

	if *sourceFile == "" {
		return errors.New("Missing file argument.")
	}

	if err := runProgram(*sourceFile); err != nil {
		return err
	}

	memoryFile, err := handleProfileFlag(*memProfile)
	if err != nil {
		return err
	}

	defer memoryFile.Close()

	runtime.GC()
	if err := pprof.WriteHeapProfile(memoryFile); err != nil {
		return errors.New("could not write memory profile: " + err.Error())
	}

	return nil
}

func runProgram(sourceFile string) error {

	tarReader, err := file.GetTgzReader(sourceFile)
	if err != nil {
		return err
	}

	if err := Indexer.IterateTarReader(tarReader); err != nil {
		return err
	}

	if err := Indexer.IndexEmailsToDatabase(); err != nil {
		return err
	}

	return nil
}

func handleProfileFlag(fileName string) (*os.File, error) {
	if fileName == "" {
		println("aaa")
		return nil, nil
	}

	f, err := file.CreateProfileFile(fileName)
	if err != nil {

		return nil, err
	}

	return f, nil
}
