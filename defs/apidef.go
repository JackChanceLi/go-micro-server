package defs

type UserIdentity struct {
	UserName string `json:"user_name"` //注册用户名
	Passwd   string `json:"password"` //注册用户的密码
	Email    string `json: "email"` //注册用户的邮箱
	Role     int    `json:"role"`  //表示用户权限，1为管理员，2位普通用户
}

type SignedUp struct{
	Success bool `json:"success"`  //session 验证是否成功
	SessionId string `json:"session_id"` //返回sessionID
}
//session处理字段
type Session struct{
	UserName string  //session对用的用户名
	TTL      int64  //session的有效期
}