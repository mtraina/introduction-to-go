package main

import (
	"path/filepath"
	"os"
	"log"
)

func main(){
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		log.Println(path)
		return nil
	})
}
