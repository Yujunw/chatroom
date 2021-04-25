package main

import (
	"chatroom/common/utils"
	"fmt"
	"net"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()

	tf := &utils.Transfer{
		Conn: conn,
	}
	err, m := tf.ReadMsg()
	if err != nil {
		fmt.Println("utils.ReadMsg failed", err)
		return
	}

	fmt.Println("读取到的消息内容为：", m)
}

func main() {
	fmt.Println("服务器在8888端口监听...")
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		// 如果监听失败，直接报错
		fmt.Println("net.Listen failed", err)
		return
	}

	defer listener.Close()

	for {
		fmt.Println("等待客户端与服务器连接...")
		conn, err := listener.Accept()
		if err != nil {
			// 一次连接失败无所谓，下一次连接可能成功
			fmt.Println("listener.Accept failed", err)
		}

		go process(conn)
	}

}
