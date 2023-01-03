package channel_demo

import (
	"fmt"
	"os"
	"time"
)

/**
*@author: 廖理
*@date:2023/1/3
**/

func PlayBall(playerName string, table chan Ball, serve bool) {
	var receive, send chan Ball
	if serve {
		receive, send = nil, table
	} else {
		receive, send = table, nil
	}
	var lastValue Ball = 1
	for {
		select {
		case send <- lastValue:
		case value := <-receive:
			fmt.Println(playerName, value)
			value += lastValue
			if value < lastValue { // 溢出了
				os.Exit(0)
			}
			lastValue = value
		}
		receive, send = send, receive // 开关切换
		time.Sleep(time.Second)
	}
}

func PlayBallDemo() {
	table := make(chan Ball)
	go PlayBall("A:", table, false)
	PlayBall("B:", table, true)
}
