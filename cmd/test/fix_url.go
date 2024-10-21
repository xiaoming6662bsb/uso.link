package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func DoFixUrl() {
	name := ""
	flag.StringVar(&name, "n", "", "")
	flag.Parse()
	if name == "" {
		return
	}
	err := filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".html") {
			err := processFile(path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", name, err)
	}
}
func processFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	updatedContent := replaceURLs(string(content))
	err = os.WriteFile(filePath, []byte(updatedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %v", filePath, err)
	}

	fmt.Printf("Processed file: %s\n", filePath)
	return nil
}

func replaceURLs(content string) string {
	re := regexp.MustCompile(`<a\s+href=\"(/[^"/]+)/?\">`)
	return re.ReplaceAllString(content, `<a href="$1.html">`)
}
