package main

import (
	"day01/internal/database"
	"day01/internal/reader"
	"flag"
	"fmt"
)

func compareDatabases(original, stolen []database.Cake) {
	origMap := make(map[string]database.Cake)
	for _, cake := range original {
		origMap[*cake.Name] = cake
	}

	stolenMap := make(map[string]database.Cake)
	for _, cake := range stolen {
		stolenMap[*cake.Name] = cake
	}

	for name := range origMap {
		if _, exists := stolenMap[name]; !exists {
			fmt.Printf("ADDED cake \"%s\"\n", name)
		}
	}

	for name := range stolenMap {
		if _, exists := origMap[name]; !exists {
			fmt.Printf("REMOVED cake \"%s\"\n", name)
		}
	}

	for name, origCake := range origMap {
		stolenCake, exists := stolenMap[name]
		if !exists {
			continue
		}

		if origCake.Time != stolenCake.Time && *origCake.Time != *stolenCake.Time {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", name, *origCake.Time, *stolenCake.Time)
		}

		origIngredientsMap := make(map[string]database.Ingredient)
		for _, ing := range origCake.Ingredients {
			origIngredientsMap[*ing.Name] = ing
		}

		stolenIngredientsMap := make(map[string]database.Ingredient)
		for _, ing := range stolenCake.Ingredients {
			stolenIngredientsMap[*ing.Name] = ing
		}

		for nameIng, origIng := range origIngredientsMap {
			stolenIng, exists := stolenIngredientsMap[nameIng]
			if !exists {
				fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", nameIng, name)
				continue
			}

			if origIng.Count != nil && stolenIng.Count != nil && *origIng.Count != *stolenIng.Count {
				fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", nameIng, name, *stolenIng.Count, *origIng.Count)
			} else if origIng.Count == nil && stolenIng.Count != nil {
				fmt.Printf("REMOVED unit count \"%s\" for ingredient \"%s\" for cake \"%s\"\n", *stolenIng.Count, nameIng, name)
			} else if origIng.Count != nil && stolenIng.Count == nil {
				fmt.Printf("ADDED unit count for ingredient \"%s\" for cake \"%s\"\n", nameIng, name)
			}

			if origIng.Unit != nil && stolenIng.Unit != nil && *origIng.Unit != *stolenIng.Unit {
				fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", nameIng, name, *stolenIng.Unit, *origIng.Unit)
			} else if origIng.Unit == nil && stolenIng.Unit != nil && *stolenIng.Unit != "" {
				fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", *stolenIng.Unit, nameIng, name)
			} else if origIng.Unit != nil && stolenIng.Unit == nil && *origIng.Unit != "" {
				fmt.Printf("ADDED unit for ingredient \"%s\" for cake \"%s\"\n", nameIng, name)
			}
		}

		for nameIng := range stolenIngredientsMap {
			if _, exists := origIngredientsMap[nameIng]; !exists {
				fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", nameIng, name)
			}
		}
	}
}

func main() {
	var originalFilePath, stolenFilePath string
	flag.StringVar(&originalFilePath, "old", "", "path to original database file (XML or JSON)")
	flag.StringVar(&stolenFilePath, "new", "", "path to stolen database file (XML or JSON)")
	flag.Parse()

	if originalFilePath == "" || stolenFilePath == "" {
		fmt.Println("ERROR | You must specify both original and stolen file paths using the --old and --new flags.")
		return
	}

	origFile := reader.FileInfo{FilePath: originalFilePath}
	_, origStruct := origFile.Read()

	stolenFile := reader.FileInfo{FilePath: stolenFilePath}
	_, stolenStruct := stolenFile.Read()

	compareDatabases(origStruct.Cakes, stolenStruct.Cakes)
}
