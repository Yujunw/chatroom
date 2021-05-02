package process

import (
	"chatroom/common/message"
	"chatroom/common/utils"
	"chatroom/server/model"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (u *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal failed", err)
		return
	}

	// 定义一个返回登录结果的消息
	var loginRepMsg message.Message
	loginRepMsg.Type = message.LoginResMsgType

	var loginResMsg message.LoginResMsg
	// 在redis数据库中去完成验证
	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOT_EXISTS {
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		} else if err == model.ERROR_USER_PWD_WRONG {
			loginResMsg.Code = 501
			loginResMsg.Error = err.Error()
		} else {
			loginResMsg.Code = 599
			loginResMsg.Error = "未知错误"
		}
		return err
	} else {
		loginResMsg.Code = 200
		fmt.Println(user, "登录成功")
	}
	//if loginMsg.UserId == 100 && loginMsg.UserPwd == "yjw" {
	//	loginResMsg.Code = 200
	//} else {
	//	loginResMsg.Code = 500
	//	loginResMsg.Error = "不存在该用户！"
	//}

	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("json.Marshal failed", err)
		return
	}

	loginRepMsg.Data = string(data)

	tf := &utils.Transfer{
		Conn: u.Conn,
	}

	err = tf.WriteMsg(data)
	if err != nil {
		fmt.Println("tf.WriteMsg failed", err)
		return
	}
	return
}
