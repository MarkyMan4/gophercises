package main

import (
	"io/fs"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func readContent(filename string) []byte {
	content, _ := os.ReadFile(filename)
	return content
}

func mdToHTML(md []byte) []byte {
	// create parser
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func main() {
	content := readContent("content/index.md")
	md := content
	html := mdToHTML(md)

	os.WriteFile("index.html", html, fs.FileMode(os.O_CREATE))
}
