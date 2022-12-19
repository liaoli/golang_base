package chat_room

import (
	"fmt"
	"net"
)

/**
*@author: 廖理
*@date:2022/12/19
**/

var clients = map[string]net.Conn{}

func MyChatRoomDemo() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("listen error ", err)
		return
	}

	for {

		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Accept error ", err)
			return
		}

		go handleClientMsg(conn)

	}

}

func handleClientMsg(conn net.Conn) {
	var romoteName string
	var buf = make([]byte, 4096)

	rAddr := conn.RemoteAddr()
	romoteName = rAddr.String()
	login := fmt.Sprintf("%s进入聊天室\n", romoteName)
	clients[romoteName] = conn

	for _, v := range clients {
		v.Write([]byte(login))
	}

	for {
		n, _ := conn.Read(buf)

		if n != 0 {
			recMsg := string(buf[:n])
			sendMsg := fmt.Sprintf("收到来自%s的消息：%s", romoteName, recMsg)
			for _, v := range clients {
				v.Write([]byte(sendMsg))
			}
		}
	}
}
