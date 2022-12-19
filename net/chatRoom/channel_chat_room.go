package chat_room

import (
	"fmt"
	"net"
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
		conn.Write([]byte(msg))
	}

}

func MessageManager() {
	onLineClients = make(map[string]*MyClient)
	for {

		msg := <-Message

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
	var romoteName string
	var buf = make([]byte, 4096)

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

	for {
		n, _ := conn.Read(buf)

		if n != 0 {
			recMsg := string(buf[:n])

			Message <- makeMsg(client, recMsg)
		}
	}
}
