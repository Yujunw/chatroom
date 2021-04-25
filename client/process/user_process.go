package process

import (
	"chatroom/common/message"
	"chatroom/common/utils"
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

	// 5. 将msg序列化，到这个时候 data就是我们要发送的消息
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
	}

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WriteMsg(data)
	if err != nil {
		fmt.Println("net.WriteMsg failed", err)
		return
	}

	//休眠10s
	time.Sleep(10 * time.Second)
}
