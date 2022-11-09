package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

/**
*@author: 廖理
*@date:2022/11/9
**/

//为什么会出现粘包
//主要原因就是tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发。
//
//“粘包”可发生在发送端也可发生在接收端：
//
//由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。
//简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，
//而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。
//接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据。
//当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。

//解决办法
//出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，
//因此我们可以对数据包进行封包和拆包的操作。
//
//封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入”包尾”内容)。
//包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。
//
//我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。

//Encode 编码生成协议
func Encode(message string) ([]byte, error) {
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)

	//写消息头
	err := binary.Write(pkg, binary.LittleEndian, length)

	if err != nil {
		return nil, err
	}
	//写消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))

	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)

	if err != nil {
		return "", err
	}

	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)

	if err != nil {
		return "", err
	}

	return string(pack[4:]), err
}
