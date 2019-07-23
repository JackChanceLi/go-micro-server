package mailop

import (
	"net/smtp"
	"strings"
	/*"math/rand"
	"time"*/
)

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")//发送端邮箱
	auth := smtp.PlainAuth("", user, password, hp[0])//用户端邮箱及其密码
	var content_type string
	if mailtype == "html" {//邮件类型
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)//发送的信息
	send_to := strings.Split(to, ";")//接收端邮箱
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

/*func Send() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	fmt.Println(vcode)
	user := "2272731163@qq.com"
	password := "j1709Y(0803)L!"
	host := "smtp.qq.com:587"
	to := "1023546080@qq.com"
	subject := "Email address verification"
	body := `<html>
			<body>
			<h3>
			"验证码: 111111"
			</h3>
			</body>
			</html>
			`
	fmt.Println("Send mail")
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("Send mail failed")
		fmt.Println("err")
	} else {
		fmt.Println("Send mail success")
	}
}*/