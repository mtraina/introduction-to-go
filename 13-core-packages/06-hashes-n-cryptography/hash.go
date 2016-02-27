package main

import (
	"hash/crc32"
	"log"
)
func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	log.Println(v)
}