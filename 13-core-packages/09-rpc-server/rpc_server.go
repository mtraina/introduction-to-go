package main

import (
	"net/rpc"
	"net"
	"log"
	"fmt"
)

type Server struct {}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server(){
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("error listening to port 9999, %v", err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

func client(){
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalf("error on dial, %v", err)
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		log.Fatalf("error calling the remote procedure, %v", err)
	} else {
		log.Println("Server.Negate(999) = ", result)
	}
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
