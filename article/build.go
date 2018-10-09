package main

import (
	"bytes"
	"io/ioutil"
	"os"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func main() {
	var options = blackfriday.WithExtensions(blackfriday.Tables | blackfriday.FencedCode)
	var articleMD, err = ioutil.ReadFile("article.md")
	if err != nil {
		panic(err)
	}
	var rendered = blackfriday.Run(articleMD, options)
	var page = bytes.NewBufferString(`<meta charset="UTF-8">`+"\n")
	page.Write(rendered)
	if err := ioutil.WriteFile("index.html", page.Bytes(), os.ModePerm); err != nil {
		panic(err)
	}
}
