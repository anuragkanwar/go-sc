package main

import (
	"fmt"
	"os"

	sitemap "github.com/oxffaa/gopher-parse-sitemap"
)

func main() {
	sitemapPath := os.Args[1]
	result := make([]string, 0)
	err := sitemap.ParseIndexFromFile(sitemapPath, func(e sitemap.IndexEntry) error {
		result = append(result, e.GetLocation())
		return nil
	})

	if err == nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
