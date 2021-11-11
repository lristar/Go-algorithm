package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
)

//go-tcpsock/read_write/server2.go
var Class map[string]map[string]int
var Data map[string]chan string

func handleConn(ctx context.Context, c net.Conn) {
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
		p := c.RemoteAddr().String()
		if strings.Contains(string(buf[:n]), "create") {
			className := AddClass(string(buf[:n]), p)
			c.Write([]byte("goto " + className))
			fmt.Println("确认", p, "的房间是:", className)
		} else if strings.Contains(string(buf[:n]), "exit") {
			exitPeople(string(buf[:n]), p)
		} else {
			data := strings.Split(string(buf[:n]), " ")
			if data[0] == "sync" {
				fmt.Println("sync")
				fmt.Println(p + ":" + string(buf[:n]))
			} else if len(data)==1 {
				fmt.Print("无人认领")
				fmt.Println(p + ":" + string(buf[:n]))
			}else {
				fmt.Print("群发")
				people := Class[data[0]]
				for e, _ := range people {
					//if v == 1 {
						Data[e] <- data[1]
						fmt.Println("向",e,"发送",data[1])
					//}
				}
				pp := <-Data[p]
				c.Write([]byte(pp))
				}
		}

	}
}
func exitPeople(read, people string) {
	className := strings.Split(read, " ")
	value, ok := Class[className[1]]
	if ok {
		value[people] = 0
		fmt.Println(people, "退出", className)
	}
}
func AddClass(read, people string) string {
	className := strings.Split(read, " ")
	v, ok := Class[className[1]]
	if !ok {
		Class[className[1]] = map[string]int{people: 1}
		fmt.Println("用户 :", people, "create :", className[1])
	} else {
		v[people] = 1
		Class[className[1]] = v
		fmt.Println("用户 :", people, "进入聊天室:", className[1])
	}
	return className[1]
}

func main() {
	ln, err := net.Listen("tcp", ":6010")
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	Class = make(map[string]map[string]int)
	Data = make(map[string]chan string)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("get client connection error: ", err)
		}
		go handleConn(ctx, conn)
	}
}
