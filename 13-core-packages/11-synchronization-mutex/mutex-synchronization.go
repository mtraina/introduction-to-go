package main

import (
	"sync"
	"log"
	"time"
	"fmt"
)

func main() {
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			log.Println(i, "start")
			time.Sleep(time.Second)
			log.Println(i, "end")
			m.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}
