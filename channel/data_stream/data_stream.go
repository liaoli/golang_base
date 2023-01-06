package data_stream

/**
*@author: 廖理
*@date:2023/1/6
**/
import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

func RandomGenerator() <-chan uint64 {
	c := make(chan uint64)
	go func() {
		rnds := make([]byte, 8)
		for {
			_, err := rand.Read(rnds)
			if err != nil {
				close(c)
				break
			}
			c <- binary.BigEndian.Uint64(rnds)
		}
	}()
	return c
}

func DataStreamDemo() {

	c := RandomGenerator()

	d := <-c

	fmt.Println("d=", d)
}
