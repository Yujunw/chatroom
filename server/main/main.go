package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

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

// 处理和客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 8096)
	// 先读取前4个字节
	_, err := conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read failed", err)
		return
	}

	var msgLen uint32
	msgLen = binary.BigEndian.Uint32(buf[:4])

	// 根据msgLen读取消息内容
	n, err := conn.Read(buf[:msgLen])
	if err != nil || n != int(msgLen) {
		fmt.Println("conn.Read failed", err)
		return
	}

	var msg message.Message
	err = json.Unmarshal(buf[:msgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal failed", err)
		return
	}

	fmt.Println("读取到的消息内容:", msg)
}
