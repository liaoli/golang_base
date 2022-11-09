package mynet

import (
	"awesomeProject/net/proto"
	"bufio"
	"fmt"
	"net"
)

/**
*@author: 廖理
*@date:2022/11/9
**/
//粘包
func process3(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {

		msg, err := proto.Decode(reader)

		if err != nil {
			fmt.Println("读取客户端发来的消息失败了")
			break
		}

		fmt.Println("客户端发来消息：", msg)

		res := "server收到了：" + msg

		data, err := proto.Encode(res)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}

		conn.Write(data)
	}
}

func StartServer3() {

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
		go process3(conn)

	}
}
