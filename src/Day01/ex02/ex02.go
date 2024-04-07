package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func readMainFile(filePath string) (map[[16]byte]string, string) {
	ans := make(map[[16]byte]string)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Sprintf("ERROR | %s", err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Sprintf("ERROR | %s", err)
		}
		value := md5.Sum(line)
		ans[value] = string(line)
	}
	return ans, ""
}
func compare(filePath string, hashMap map[[16]byte]string) (map[[16]byte]string, string) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Sprintf("ERROR | %s", err)
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Sprintf("ERROR | %s", err)
		}
		value := md5.Sum(line)
		if _, exists := hashMap[value]; !exists {
			fmt.Printf("ADDED %s\n", line)
		} else {
			delete(hashMap, value)
		}
	}
	return hashMap, ""
}

func main() {
	var origSnapPath, newSnapPath string
	flag.StringVar(&origSnapPath, "old", "", "path to original database file (XML or JSON)")
	flag.StringVar(&newSnapPath, "new", "", "path to stolen database file (XML or JSON)")
	flag.Parse()

	if origSnapPath == "" || newSnapPath == "" {
		fmt.Println("ERROR | You must specify both original and stolen file paths using the --old and --new flags.")
		return
	}

	if !(strings.HasSuffix(origSnapPath, ".txt") && strings.HasSuffix(newSnapPath, ".txt")) {
		fmt.Println("ERROR | Program working only with .txt files")
		return
	}

	hashMap, err := readMainFile(origSnapPath)
	if err != "" {
		fmt.Println(err)
		return
	}
	hashMap, err = compare(newSnapPath, hashMap)
	if err != "" {
		fmt.Println(err)
		return
	}
	for _, value := range hashMap {
		fmt.Printf("REMOVED %s\n", value)
	}
}
