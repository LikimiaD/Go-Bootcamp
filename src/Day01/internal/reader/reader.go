package reader

import (
	"bufio"
	"day01/internal/database"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type DBReader interface {
	Read() (string, database.Recipes)
}

type FileInfo struct {
	FilePath string
}

func (f *FileInfo) Read() (string, database.Recipes) {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return "ERROR | The program cannot open the file, it may not exist.", database.Recipes{}
	}
	defer file.Close()

	r := bufio.NewReader(file)
	if strings.HasSuffix(f.FilePath, ".json") {
		return f.readJSON(r)
	} else if strings.HasSuffix(f.FilePath, ".xml") {
		return f.readXML(r)
	}
	return "ERROR | Unsupported file format.", database.Recipes{}
}

func (f *FileInfo) readJSON(r *bufio.Reader) (string, database.Recipes) {
	dec := json.NewDecoder(r)
	var res database.Recipes
	if err := dec.Decode(&res); err != nil {
		return fmt.Sprintf("ERROR | Error decoding JSON: %s", err), database.Recipes{}
	}
	b, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		return fmt.Sprintf("ERROR | Error marshalling JSON: %s", err), database.Recipes{}
	}
	return string(b), res
}

func (f *FileInfo) readXML(r *bufio.Reader) (string, database.Recipes) {
	dec := xml.NewDecoder(r)
	var res database.Recipes
	if err := dec.Decode(&res); err != nil {
		return fmt.Sprintf("Error decoding XML: %s", err), database.Recipes{}
	}
	b, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		return fmt.Sprintf("Error marshalling XML: %s", err), database.Recipes{}
	}
	return string(b), res
}
