package main

import (
	"net"
	"log"
	"encoding/gob"
	"fmt"
)

func server() {
	// listen to a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("error listening to port 9999: %v", err)
	}
	for {
		// accept connection
		c, err := ln.Accept()
		if err != nil {
			log.Fatalf("error accepting connection: %v", err)
		}
		// handle the connection
		go handleServerConnection(c)
		log.Print("server started")
	}
}

func handleServerConnection(c net.Conn){
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		log.Fatalf("error decoding the message: %v", err)
	} else {
		log.Println("received: %v", msg)
	}
	c.Close()
}

func client(){
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalf("error connecting to the server: %v", err)
	}

	// send the message
	msg := "Hello World"
	log.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		log.Fatalf("error encoding the message: %v", err)
	}
	c.Close()
}

func main(){
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
