package main

import (
	"encoding/xml"
	"log"
	"os"
	"path/filepath"
)

var cwd, err = os.Getwd()

func main() {

	if err != nil {
		log.Fatal("Error getting file directory", err)
	}
	files := os.Args[1:]
	if len(files) < 1 {
		files = readCwd(cwd)
	}

	files = filterFiles(files)

	items := generateItems(files)

	startTime := StartTime{
		Year:   2020,
		Month:  03,
		Hour:   00,
		Minute: 00,
		Second: 00,
	}

	background := Background{
		StartTime: startTime,
		Comment:   "This animation will start at midnight.",
		Items:     items,
	}
	if len(items) < 1 {
		log.Fatal("Cannot create XML config, check that you passed valid files")
	}
	xmlFile, err := os.Create(filepath.Base(cwd) + ".xml")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer xmlFile.Close()
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&background)
	if err != nil {
		log.Fatal("Error marshalling to XML:", err)
	}
}
