package message

// 在此定义发送消息的类型
// 1. 基本类型的消息，真正发送的消息
// 2. 登录消息，客户端发送给服务器
// 3. 登录结果，服务器发送给客户端

const ()

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
