package main

import (
	"io/ioutil"
	"log"
)

func main(){
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatalf("error reading the file %v", err)
	}
	str := string(bs)
	log.Println(str)
}
