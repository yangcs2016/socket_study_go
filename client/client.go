package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "10.28.101.106:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Printf("conn success conn=%v\n", conn)
}
