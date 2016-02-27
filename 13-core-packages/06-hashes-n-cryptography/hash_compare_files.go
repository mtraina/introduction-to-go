package main

import (
	"hash/crc32"
	"io/ioutil"
	"log"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("13-core-packages/06-hashes-n-cryptography/test1.txt")
	if err != nil {
		log.Fatalf("error getting the hash of the file test1.txt %v", err)
	}
	h2, err := getHash("13-core-packages/06-hashes-n-cryptography/test2.txt")
	if err != nil {
		log.Fatalf("error getting the hash of the file test2.txt %v", err)
	}
	log.Println(h1, h2, h1 == h2)
}