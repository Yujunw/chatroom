package main

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	processPkg "chatroom/server/process"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (p *Processor) serverProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		up := processPkg.UserProcess{
			Conn: p.Conn,
		}
		err := up.ServerProcessLogin(msg)
		if err != nil {
			return err
		}
	case message.RegisterMsgType:
		fmt.Println("注册消息")
	default:
		fmt.Println("不支持的消息类别！")
	}

	return
}

// 循环处理客户端发送的消息
func (p *Processor) loopProcessMsg() (err error) {
	for {
		tf := utils.Transfer{
			Conn: p.Conn,
		}
		msg, err := tf.ReadMsg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出..")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}

		err = p.serverProcessMsg(&msg)
		if err != nil {
			fmt.Println("serverProcessMsg failed ", err)
			return err
		}
	}
	return
}
