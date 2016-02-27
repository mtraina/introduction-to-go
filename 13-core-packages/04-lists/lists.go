package main

import (
	"container/list"
	"log"
)

func main() {
	var x list.List

	// append to the list
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		log.Println(e.Value.(int))
	}
}
