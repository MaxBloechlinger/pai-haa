package main

import (
	"fmt"

	"gosearch/index"
)

func main() {
	docs := []string{
		"Go makes search engines",
		"Search with inverted index in Go",
		"Crawlers feed the index",
	}
	idx := index.New(docs)
	fmt.Println(idx.Search("go"))
}
