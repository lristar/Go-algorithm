package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

//func main() {
//	conn, err := net.Dial("tcp", ":6010")
//	if err != nil {
//		panic(err)
//	}
//	data, err := bufio.NewReader(conn).ReadString('\n')
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("%#v\n", data)
//}


func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: go run client2.go YOUR_CONTENT")
		return
	}
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":6010")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok")

	time.Sleep(time.Second * 2)
	data := os.Args[1]
	conn.Write([]byte(data))

	for   {
		input:=bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text()=="exit"{
			break
		}
		conn.Write([]byte(input.Text()))
		fmt.Println("发送",input.Text())

	}
}
