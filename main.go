package main

import (
	"Indexer/Indexer"
	"Indexer/file"
	"errors"
	"flag"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memProfile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	flag.Parse()
	if err := initializeProgram(); err != nil {
		os.Exit(1)
	}
}

func initializeProgram() error {

	cpuFile, err := handleProfileFlag(*cpuProfile)
	if err != nil {
		return err
	}
	cpuFile.Close()

	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		return errors.New("could not start CPU profile: " + err.Error())
	}
	defer pprof.StopCPUProfile()

	sourceFile := flag.Arg(0)

	if sourceFile == "" {
		return errors.New("Missing file argument.")
	}

	if err := runProgram(sourceFile); err != nil {
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
		return nil, nil
	}

	f, err := file.CreateProfileFile(fileName)
	if err != nil {
		return nil, err
	}

	return f, nil
}
