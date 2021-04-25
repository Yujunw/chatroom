package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (t *Transfer) WriteMsg(data []byte) (err error) {
	// 6.1 先把 data的长度发送给服务器
	// 先获取到 data的长度->转成一个表示长度的byte切片
	var msgLen uint32
	msgLen = uint32(len(data))
	// 大端序
	binary.BigEndian.PutUint32(t.Buf[:4], msgLen)
	n, err := t.Conn.Write(t.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err =", err)
		return
	}

	// 6.2 发送消息本身
	n, err = t.Conn.Write(data)
	if n != int(msgLen) || err != nil {
		fmt.Println("conn.Write err =", err)
		return
	}
	fmt.Printf("客户端发送消息的长度=%d 内容=%s\n", len(data), string(data))
	return
}

func (t *Transfer) ReadMsg() (err error, msg message.Message) {
	buf := make([]byte, 8096)
	// 先读取前4个字节
	_, err = t.Conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read failed", err)
		return
	}

	var msgLen uint32
	msgLen = binary.BigEndian.Uint32(buf[:4])

	// 根据msgLen读取消息内容
	n, err := t.Conn.Read(buf[:msgLen])
	if err != nil || n != int(msgLen) {
		fmt.Println("conn.Read failed", err)
		return
	}
	err = json.Unmarshal(buf[:msgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal failed", err)
		return
	}

	return
}
