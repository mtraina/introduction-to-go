package main

import (
	"os"
	"log"
)

func main(){
	file, err := os.Create("13-core-packages/02-create-files/test.txt")
	if err != nil {
		log.Fatalf("error creating the file %v", err)
	}
	defer file.Close()

	file.WriteString("test")
	log.Println("end")
}
