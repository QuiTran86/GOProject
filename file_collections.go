package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func ScanDirectory(path string) error {
	fmt.Println("Scanning path:", path, "...")
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := ScanDirectory(filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil

}

func main() {
	err := ScanDirectory(".")
	if err != nil {
		log.Fatal(err)
	}
}
