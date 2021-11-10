package main

import (
	"fmt"
	"log"
	"net"
)
//go-tcpsock/read_write/server2.go

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		var buf = make([]byte, 1024*1024)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Println("conn read error:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}


func main() {
	ln, err := net.Listen("tcp", ":6010")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("get client connection error: ", err)
		}
		go handleConn(conn)
	}
}