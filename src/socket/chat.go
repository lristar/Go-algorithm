package main

import (
"bufio"
"fmt"
"log"
"net"
"strings"
"time"
)

var globalRoom *Room = NewRoom()

type Room struct {
	users map[string]net.Conn
}

func NewRoom() *Room {
	return &Room{
		users: make(map[string]net.Conn),
	}
}

func (r *Room) Join(user string, conn net.Conn) {
	_, ok := r.users[user]
	if ok {
		r.Leave(user) //如果存在用户user就提出之前的链接，调用Leave方法实现。
	}
	r.users[user] = conn
	fmt.Printf("%s 登录成功。\n", user)
	conn.Write([]byte(user + ":加入聊天室!\n"))
}

func (r *Room) Leave(user string) {
	conn, ok := r.users[user]
	if !ok {
		fmt.Printf("%v用户不存在！", user)
	}
	conn.Close() //如果存在用户就断开链接。
	delete(r.users, user) //将用户从字典中删除。
	fmt.Printf("%s 离开", user)
}

func (r *Room) Broadcast(who string, msg string) {
	time_info := time.Now().Format("2006年01月02日 15:04:05") //这个是对日期定义一个格式化输出。告诉你一个记住它的方法：2006-01-02 15:04:05对应着2006 1(01) 2(02) 3(15) 4(04) 5(05) 哈哈
	tosend := fmt.Sprintf("%v %s:%s\n", time_info,who, msg)
	for user, conn := range r.users { //遍历所有用户，
		if user == who {
			continue //当发现用户是自己就不发送数据。即跳过循环。
		}
		conn.Write([]byte(tosend)) //将数据发送给登陆的用户。
	}
}

func Handle_Conn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("read:")
	r := bufio.NewReader(conn) //将用户的输入存入“r”中，方便一会我们按块读取。
	line, err := r.ReadString('\n')
	fmt.Println("line:",line)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("aaa")
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields) != 2 {
		conn.Write([]byte("您输入的字符串用户名活密码无效，程序强制退出！\n"))
		return
	}
	fmt.Println("bbb")
	user := fields[0]
	password := fields[1]
	if password != "123" {
		return
	}
	globalRoom.Join(user, conn)
	globalRoom.Broadcast("System", fmt.Sprintf("%s join room", user))
	for { //获取用户的输入。
		conn.Write([]byte("按回车键发送消息：>>>"))//这里是给客户端增加一个提示符
		line, err := r.ReadString('\n') //循环读取用户输入的内容。换行符为“\n”
		if err != nil {                 //当用户主动关闭连接是，会出现报错就直接直接终止循环。
			break
		}
		line = strings.TrimSpace(line)   //去掉换行符
		fmt.Println(user,line)
		globalRoom.Broadcast(user, line) // 将用户输入的消息进行广播。
	}
	globalRoom.Broadcast("System", fmt.Sprintf("%s Leave room", user))
	globalRoom.Leave(user) //踢掉用户。
}

func main() {
	addr := "0.0.0.0:8888"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("远程链接：",conn.RemoteAddr())
		go Handle_Conn(conn)
	}
}