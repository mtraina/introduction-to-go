package main

import (
	"os"
	"log"
)

func main(){
	dir, err := os.Open(".")
	if err != nil {
		log.Fatalf("error opening directory %v", err)
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		log.Fatalf("error reading directory content %v", err)
	}

	for _, fi := range fileInfos {
		log.Println(fi.Name())
	}
}
