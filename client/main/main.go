package main

import (
	"chatroom/client/process"
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id, 一个表示用户密码
var userId int
var userPwd string

func main() {
	//接收用户的选择
	var key int
	fmt.Println("----------------欢迎登陆多人聊天系统------------")
	fmt.Println("\t\t\t 1 登陆聊天室")
	fmt.Println("\t\t\t 2 注册用户")
	fmt.Println("\t\t\t 3 退出系统")
	fmt.Println("\t\t\t 请选择(1-3):")

	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("登陆聊天室")
	case 2:
		fmt.Println("注册用户")
	case 3:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("你的输入有误，请重新输入")
	}
	//更加用户的输入，显示新的提示信息
	if key == 1 {
		//说明用户要登陆
		fmt.Println("请输入用户的id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户的密码")
		fmt.Scanf("%s\n", &userPwd)
		//先把登陆的函数，写到另外一个文件， 比如login.go
		process.Login(userId, userPwd)
	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑....")
	} else {
		os.Exit(0)
	}
}
