package udp

import (
	"fmt"
	"net"
)

/**
*@author: 廖理
*@date:2022/11/9
**/

func StartClient() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})

	defer socket.Close()
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}

	sendData := []byte("hello server")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("发送数据给服务端失败，err:", err)
		return
	}
	recData := make([]byte, 4096)
	n, addr, err := socket.ReadFromUDP(recData)

	if err != nil {
		fmt.Println("接收服务端数据失败，err:", err)
		return
	}

	fmt.Printf("recv:%v addr:%v count:%v\n", string(recData[:n]), addr, n)
}
