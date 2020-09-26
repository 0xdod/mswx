package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
)

func filterFiles(files []string) []string {
	var newFiles []string
	for _, file := range files {
		if !isExtAllowed(path.Ext(file)) {
			fmt.Printf("skipping %s: invalid file or bad extension\n", filepath.Base(file))
			continue
		}
		newFiles = append(newFiles, file)
	}
	return newFiles
}

func isExtAllowed(ext string) bool {
	allowedExt := []string{".png", ".jpg", ".jpeg", ".gif", ".webp", ".svg"}
	for _, el := range allowedExt {
		if ext == el {
			return true
		}
	}
	return false
}

func readCwd(cwd string) []string {
	var files []string
	entries, err := ioutil.ReadDir(".")
	if err == nil {
		for _, file := range entries {
			if file.IsDir() {
				continue
			}
			files = append(files, file.Name())
		}
	}
	return files
}

func generateItems(files []string) []interface{} {
	var items []interface{}

	for i, file := range files {
		if !isExtAllowed(path.Ext(file)) {
			log.Printf("skipping %s: invalid file bad extension", file)
			continue
		}

		abs := path.Join(cwd, file)
		tr := Transition{
			Duration: 3.00,
			From:     abs,
		}
		if i == len(files)-1 {
			tr.To = path.Join(cwd, files[0])
		} else {
			tr.To = path.Join(cwd, files[i+1])
		}
		static := Static{
			Duration: 120,
			File:     abs,
		}
		items = append(items, &static)
		items = append(items, &tr)
	}
	return items
}
