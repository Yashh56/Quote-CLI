package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type Quote struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func main() {

	res, err := http.Get("https://stoic.tekloon.net/stoic-quote")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		color.Red("Api is not available !!")
	}

	var quoteRes Quote

	err = json.Unmarshal(body, &quoteRes)

	if err != nil {
		log.Fatal(err)
	}

	author, quote := quoteRes.Author, quoteRes.Quote

	if author != " " && quote != " " {
		color.Cyan(quote)
		fmt.Println()
		fmt.Print("-")
		color.Cyan(author)
	}

}
