package model

type User struct {
	// 为了序列化和反序列化成功，必须保证json字符串的key和结构体字段的tag名字一致
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
