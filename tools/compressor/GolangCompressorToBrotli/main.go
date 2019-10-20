package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/kothar/brotli-go.v0/enc"
)

var file = flag.String("file", "", "Path to file to compress")

func main() {
	flag.Parse()
	fmt.Println(*file)
	if *file == "" {
		log.Fatal("You didn't pass file to compress")
	}
	fileData, err := ioutil.ReadFile(*file)
	fmt.Println("Reading data from file...")
	if err != nil {
		log.Fatal("Cannot read data from the file")
	}
	fmt.Println("Starting compressing...")
	compressedData, err := enc.CompressBuffer(nil, fileData, make([]byte, 0))
	if err != nil {
		log.Fatal("Cannot compress data")
	}
	fmt.Println("The compression has ended")
	fmt.Println("Creating new output file...")
	newFile, err := os.Create(*file + ".br")
	if err != nil {
		log.Fatal("Cannot create new file with output data")
	}
	defer newFile.Close()
	fmt.Println("Writing to file...")
	_, err = newFile.Write(compressedData)
	if err != nil {
		log.Fatal("Cannot write data to file")
	}
	fmt.Println("Data has been saved successfully")
}
