package main

import (
	"io/ioutil"
	"log"
)

func main(){
	bs, err := ioutil.ReadFile("13-core-packages/01-read-files/test.txt")
	if err != nil {
		log.Fatalf("error reading the file %v", err)
	}
	str := string(bs)
	log.Println(str)
}
