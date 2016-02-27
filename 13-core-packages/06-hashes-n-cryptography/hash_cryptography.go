package main

import (
	"crypto/sha1"
	"log"
)

func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	log.Println(bs)
}