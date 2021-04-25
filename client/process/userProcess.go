package process

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func Login(userId int, userPwd string) {
	// 1. 连接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}

	defer conn.Close()

	// 2. 通过conn发送消息，只发送通用类型消息
	var msg message.Message
	msg.Type = message.LoginMsgType

	// 3. 定义一个LoginMsg类型的消息
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	// 4. 将LoginMsg序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	msg.Data = string(data)

	// 5. 将msg序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
	}

	// 到这个时候 data就是我们要发送的消息
	// 6.1 先把 data的长度发送给服务器
	// 先获取到 data的长度->转成一个表示长度的byte切片
	var msgLen uint32
	msgLen = uint32(len(data))
	var buf [4]byte
	// 大端序
	binary.BigEndian.PutUint32(buf[:4], msgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err =", err)
	}

	// 6.2 发送消息本身
	n, err = conn.Write(data)
	if n != int(msgLen) || err != nil {
		fmt.Println("conn.Write err =", err)
	}

	fmt.Printf("客户端发送消息的长度=%d 内容=%s\n", len(data), string(data))

	//休眠10s
	time.Sleep(10 * time.Second)
}
