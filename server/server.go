package main

import (
	"fmt"
	"net" //做socket开发，net包含所有需要的方法和函数
)

//服务器端
//1.监听8888端口
func main() {
	fmt.Println("服务器开始监听...")
	//net.Listen("tcp","0.0.0.0:8888")
	//1.tcp 表示使用网络协议是tcp
	//2.0.0.0.0:8888 表示监听本地8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭listen
	//循环等待客户端来连接我
	for {
		//等待客户端来连接
		fmt.Println("等待客户端连接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept() err=", err)
		} else {
			fmt.Printf("Accept() success conn=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//可以用telnet 127.0.0.1:8888测试是否能连接
		//这里准备起一个协程，为客户端服务
	}

	fmt.Printf("listen success=%v\n", listen)
}
