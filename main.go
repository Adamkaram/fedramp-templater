package main

import (
	"fmt"
	"os"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"github.com/opencontrol/doc-template/docx"
)

func parseArgs() (inputPath, outputPath string) {
	if len(os.Args) != 3 {
		fmt.Fprint(os.Stderr, "Usage:\n\n\tfedramp-templater <input> <output>\n\n")
		os.Exit(1)
	}
	inputPath = os.Args[1]
	outputPath = os.Args[2]
	return
}

func getXmlDoc(path string) (xmlDoc *xml.XmlDocument, err error) {
	wordDoc := new(docx.Docx)
	wordDoc.ReadFile(path)

	content := wordDoc.GetContent()
	// http://stackoverflow.com/a/28261008/358804
	bytes := []byte(content)

	xmlDoc, err = gokogiri.ParseXml(bytes)
	return
}

func handleXmlDoc(doc *xml.XmlDocument) (err error) {
	// http://stackoverflow.com/a/27475227/358804
	xp := doc.DocXPathCtx()
	xp.RegisterNamespace("w", "http://schemas.openxmlformats.org/wordprocessingml/2006/main")

	query := xpath.Compile("//w:tbl")
	nodes, err := doc.Search(query)
	if err != nil {
		return
	}
	for _, node := range nodes {
		fmt.Printf("%#v\n", node)
	}

	return
}

func main() {
	// inputPath, outputPath := parseArgs()
	inputPath, _ := parseArgs()

	doc, err := getXmlDoc(inputPath)
	defer doc.Free()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = handleXmlDoc(doc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// doc.WriteToFile(outputPath, content)
}