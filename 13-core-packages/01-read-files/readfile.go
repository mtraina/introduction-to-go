package main

import (
	"os"
	"fmt"
	"log"
)

func main(){
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("error opening the file %v", err)
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("error getting the file size %v", err)
		return
	}

	// read file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatalf("error reading the file %v", err)
		return
	}

	str := string(bs)
	fmt.Println(str)
}
