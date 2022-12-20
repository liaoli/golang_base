package chat_room

import (
	"fmt"
	"net"
	"time"
)

/**
*@author: 廖理
*@date:2022/12/19
**/

var onLineClients map[string]*MyClient

type MyClient struct {
	name string
	addr string
	C    chan string
}

var Message = make(chan string)

func makeMsg(c *MyClient, text string) (msg string) {
	msg = fmt.Sprintf("%s:%s", c.name, text)
	return
}

func writeMsgToClient(client *MyClient, conn net.Conn) {
	for msg := range client.C {
		conn.Write([]byte(msg + "\n"))
	}

}

func MessageManager() {
	onLineClients = make(map[string]*MyClient)
	for {

		msg := <-Message

		fmt.Println(msg)
		for _, v := range onLineClients {
			v.C <- msg
		}
	}
}

func ChannelChatRoomDemo() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("listen error ", err)
		return
	}

	defer listener.Close()
	go MessageManager()

	for {

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Accept error ", err)
			return
		}

		go handleConn(conn)

	}

}

func handleConn(conn net.Conn) {

	defer conn.Close()

	hasData := make(chan bool)
	isQuit := make(chan bool)
	var romoteName string
	rAddr := conn.RemoteAddr()
	romoteName = rAddr.String()

	client := &MyClient{
		name: romoteName,
		addr: romoteName,
		C:    make(chan string),
	}
	onLineClients[romoteName] = client

	go writeMsgToClient(client, conn)

	Message <- makeMsg(client, "login")

	go func() {
		var buf = make([]byte, 4096)
		for {
			n, _ := conn.Read(buf)

			if n != 0 {
				recMsg := string(buf[:n-1])
				switch recMsg {
				case "who":
					//获取在线用户列表
					conn.Write([]byte("在线用户列表：\n"))
					for _, c := range onLineClients {
						conn.Write([]byte(c.name + "\n"))
					}
				default:
					Message <- makeMsg(client, recMsg)
				}
				hasData <- true
			} else {
				isQuit <- true
				return
			}

		}
	}()

	for {
		select {
		case <-isQuit:
			delete(onLineClients, client.name)
			Message <- makeMsg(client, "退出")
			return
		case <-hasData:
		//有数据啥也不做
		case <-time.After(time.Second * 60):
			//60分钟不说话踢出
			delete(onLineClients, client.name)
			Message <- makeMsg(client, "超过60秒不说话，下线")
			return
		}
	}

}
