package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	sitemap "github.com/oxffaa/gopher-parse-sitemap"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	sitemapPath := os.Args[1]
	filePath := os.Args[2]
	fmt.Println(sitemapPath)
	result := make([]string, 0)
	err := sitemap.ParseFromSite(sitemapPath, func(e sitemap.Entry) error {
		result = append(result, e.GetLocation())
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	} else {
		fmt.Println(result)
	}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	person1 := &Person{}
	byteValye, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(byteValye, person1)

	fmt.Println(person1.Name, person1.Age)
}
