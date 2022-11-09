package tcp

import (
	proto2 "awesomeProject/net/tcp/proto"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
*@author: 廖理
*@date:2022/11/9
**/
func startClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
	}

	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n') //读取用户数据
		inputInfo := strings.Trim(input, "\r\n")

		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err = conn.Write([]byte(inputInfo)) //发送消息
		if err != nil {
			fmt.Println("发送消息失败，err:", err)
			return
		}
		buf := [512]byte{}

		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("接收服务端消息失败，err:", err)
			return
		}
		fmt.Println("接收到服务端消息:", string(buf[:n]))
	}
}

func startClient2() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
	}

	defer conn.Close()

	//inputReader := bufio.NewReader(os.Stdin)

	for i := 0; i < 20; i++ {
		conn.Write([]byte("hello,你好，吃了吗!"))
	}

	buf := [512]byte{}

	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("接收服务端消息失败，err:", err)
		return
	}
	fmt.Println("接收到服务端消息:", string(buf[:n]))

}

func startClient3() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
	}

	defer conn.Close()

	//inputReader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto2.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}

	reader := bufio.NewReader(conn)
	for {
		msg, err := proto2.Decode(reader)

		if err != nil {
			fmt.Println("接收服务端消息失败，err:", err)
			return
		}
		fmt.Println("接收到服务端消息:", msg)
	}

}

func startClient4() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
	}

	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(conn)
	for {

		input, _ := inputReader.ReadString('\n') //读取用户数据
		inputInfo := strings.Trim(input, "\r\n")

		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		data, err := proto2.Encode(inputInfo)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)

		msg, err := proto2.Decode(reader)

		if err != nil {
			fmt.Println("接收服务端消息失败，err:", err)
			return
		}
		fmt.Println("接收到服务端消息:", msg)
	}

}
