package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// UserDao 完成对User的各种操作
type UserDao struct {
	// redis连接池
	pool *redis.Pool
}

// 在服务器启动之后，就需要初始化一个UserDao实例
// 将其作为全局变量，在需要redis操作时，直接使用即可
var (
	MyUserDao *UserDao
)

// 使用工厂模式，创建user
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// UserDao应该具备的功能：根据用户id在redis中查询用户信息
func (d *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil {
		// 表示未找到对应ID
		if err == redis.ErrNil {
			err = ERROR_USER_NOT_EXISTS
		}
		return
	}

	// todo 重新初始化一个user ???
	user = &User{}
	// 将res反序列化成User实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal failed ", err)
		return
	}
	return
}

// 完成对登录的校验
// 如果用户id和密码都正确，返回一个user实例，否则err
func (u *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	// 从redis连接池获取一条连接
	conn := u.pool.Get()
	defer conn.Close()

	user, err = u.getUserById(conn, userId)
	if err != nil {
		return
	}

	// 此时只能保证获取到用户，还需要判断密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD_WRONG
		return
	}
	return
}
