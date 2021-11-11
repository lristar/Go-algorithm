package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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

var class string

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
		go Read(conn)
		input:=bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text()=="exit"{
			break
		}
		if input.Text()=="query"{
			fmt.Println("class is "+class)
			continue
		}
		if input.Text()=="local"{
			fmt.Println("local is "+conn.LocalAddr().String())
			continue
		}
		if class !=""{
			conn.Write([]byte(class+" "+input.Text()))
			fmt.Println("class:",class,input.Text())
		}else {
			conn.Write([]byte(input.Text()))
			fmt.Println("发送",input.Text())
		}

	}
}

func Read(conn net.Conn){
	var buf = make([]byte, 1024*1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("conn read error:", err)
		return
	}
	fmt.Println("buff:",string(buf[:n]))
	content:=string(buf[:n])
	if strings.Contains(content,"goto"){
		data:=strings.Split(content," ")
		class=data[1]
		fmt.Println("进入了:",class)
	}

}