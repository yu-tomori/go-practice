package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk("dir",
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".go" {
				fmt.Println(path)
			}
			return nil
		})
	if err != nil {
		os.Exit(1)
	}
}
