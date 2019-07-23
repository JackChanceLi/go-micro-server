package mailop

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	//dbClear()
	m.Run()
	//dbClear()
}

func TestUserWorkFlow(t *testing.T) {

	t.Run("retrieve_password", testSendToMailRetrievePassword)
	t.Run("register", testSendToMailRegister)
}

func testSendToMailRegister(t *testing.T) {//注册发送验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := rnd.Int31n(1000000)
	user := "2272731163@qq.com"
	password := "flxkivkbilrfdijd"
	host := "smtp.qq.com:25"
	to := "1023546080@qq.com"
	subject := "Email address verification"
	body := strconv.Itoa(int(vcode))
	fmt.Println("Send mail")
	fmt.Println("Verification code:")
	fmt.Println(vcode)
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		t.Errorf("Send mail retrieve failed\nerr:%v", err)
	} else {
		t.Errorf("Send mail retrieve success")
	}
}

func testSendToMailRetrievePassword(t *testing.T) {//找回密码发送链接验证
	user := "2272731163@qq.com"
	password := "flxkivkbilrfdijd"
	host := "smtp.qq.com:25"
	to := "1023546080@qq.com"
	subject := "Email address verification"
	body := `<html>
			<body>
			<a href="http://www.baidu.com" title="Baidu">找回密码</a>
			</body>
			</html>
			`
	err := SendToMail(user, password, host, to, subject, body, "html")
	if err != nil {
		t.Errorf("Send mail retrieve failed\nerr:%v", err)
	} else {
		t.Errorf("Send mail retrieve success")
	}
}