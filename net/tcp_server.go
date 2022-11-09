package mynet

import (
	"bufio"
	"fmt"
	"net"
)

/**
*@author: 廖理
*@date:2022/11/9
**/

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)

		var buf [128]byte

		n, err := reader.Read(buf[:])

		if err != nil {
			fmt.Println("读取客户端发来的消息失败了")
			break
		}

		recStr := string(buf[:n])
		fmt.Println("客户端发来消息：", recStr)

		conn.Write([]byte("server收到了：" + recStr))
	}
}

func StartServer() {

	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err：", err)
		return
	}

	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("建立连接失败")
			continue
		}
		go process(conn)

	}
}
