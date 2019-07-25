package defs

type UserIdentity struct {
	UserName string `json:"user_name"` //注册用户名
	Password   string `json:"password"` //注册用户的密码
	Email    string `json: "email"` //注册用户的邮箱
	//Role    int  `json:"role"`  //表示用户权限，1为管理员，2位普通用户
}

type SignedUp struct{
	Success bool `json:"success"`  //session 验证是否成功
	SessionId string `json:"session_id"` //返回sessionID
	Cid string `json:id`
	Name string `json:"name"`
	Email string `json:"email"`
	Auth string `json:"auth"`

}

type UserInformation struct {
	Cid string `json:id`
	Name string `json:"name"`
	Email string `json:"email"`
	Auth string `json:"auth"`
}

type Register struct {
	Success bool `json:"success"`
	Username string `json:"user_name"`
}
//session处理字段
type Session struct{
	Uid string  //session对用的用户名
	TTL      int64  //session的有效期
}

type LiveRoomIdentity struct {
	Lid string `json:"lid"`
	Name string  `json:"name"`
	Kind string `json:"kind"`
	Size string `json:"size"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	PushUrl string `json:"push_url"`
	PullHlsUrl string `json:"pull_hls_url"`
	PullRtmpUrl string `json:"pull_rtmp_url"`
	PullHttpFlvUrl string `json:"pull_http_flv_url"`
	DisplayUrl string `json:"display_url"`
	Status string `json:"status"`
	Permission string `json:"permission"`
}


var (
	EmptyUser = UserInformation{Cid:"", Name:"", Email:"", Auth:""}
	EmptySignedUp = SignedUp{Success:false, SessionId:"", Cid:"", Name:"", Email:"", Auth:""}
)