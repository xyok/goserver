package main

import (
	"net"
	"fmt"
	"log"
	"os"
	"sync"
	"strconv"
)

//发送信息
func sender(conn net.Conn, g int) {
	defer conn.Close()
	for i := 0; i < 500; i++ {

		words := "client:[" + strconv.Itoa(g) + "]" + strconv.Itoa(i) + "Hello Server!\n"
		conn.Write([]byte(words))

	}
	fmt.Println("send over")
}

//日志
func Log(v ...interface{}) {
	log.Println(v...)
}

func main() {
	server := "127.0.0.1:3333"
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(n int) {
			defer wg.Done()
			tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
				os.Exit(1)
			}
			conn, err := net.DialTCP("tcp", nil, tcpAddr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
				os.Exit(1)
			}

			fmt.Println("connection success")
			sender(conn, n)
		}(i)

	}
	wg.Wait()

}
