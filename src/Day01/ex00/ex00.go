package main

import (
	"day01/internal/reader"
	"flag"
	"fmt"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "f", "", "path to database file (XML or JSON)")
	flag.Parse()

	if filePath == "" {
		fmt.Println("ERROR | You must specify a file path using the -f flag.")
		return
	}

	var fileInfo reader.DBReader = &reader.FileInfo{FilePath: filePath}
	output, _ := fileInfo.Read()
	fmt.Println(output)
}
