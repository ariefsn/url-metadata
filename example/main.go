package main

import (
	"encoding/json"
	"fmt"

	"github.com/ariefsn/urlmetadata"
)

func main() {
	url := "https://github.com/ariefsn"

	opt := urlmetadata.UrlMetadataParseOptions{
		Body: false,
	}
	meta, err := urlmetadata.Parse(url, opt)
	if err != nil {
		fmt.Println("Err", err)
	}

	b, _ := json.MarshalIndent(meta, "", "  ")

	fmt.Println(string(b))
}
