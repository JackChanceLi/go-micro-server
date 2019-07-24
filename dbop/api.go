package dbop

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-micro-server/defs"
	"go-micro-server/utils"
	"log"
	"time"
)

func check(err error){
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}
//获取当前的时间
func getCurrentTime() (string, error) {
	tNow := time.Now()
	timeNow := tNow.Format("2006-01-02 15:04:05")
	return timeNow, nil
}

func UserLogin(email string) (*defs.UserInformation, string, error) {
	actOut,err := dbConn.Prepare("SELECT cid, uname, password, auth from company WHERE email = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, "", err
	}
	var cid, uname, auth, password string
	err = actOut.QueryRow(email).Scan(&cid, &uname, &password, &auth)
	if err != nil && err != sql.ErrNoRows {
		return nil, "", err
	}
	defer actOut.Close()
	UI := &defs.UserInformation{}
	UI.Cid = cid
	UI.Name = uname
	UI.Auth = auth
	UI.Email = email

	return UI, password, nil
}
func userRegister(userName string, email string, password string, role int) error {
	uid, _ := utils.NewUUID()
	log.Printf("uid:%s",uid)
	actIns,err := dbConn.Prepare("INSERT  INTO user_information (user_id, user_name, register_date, email, user_passwd, role) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
    tNow, _ := getCurrentTime()
	_, err = actIns.Exec(uid, userName, tNow, email, password, role)
	if err != nil {
		return err
	}
	defer actIns.Close()
	return nil
}

func UserRegister(uname, email, password string) error {
	cid, _ := utils.NewUUID()
	log.Printf("User Register uid:%s",cid)
	actIns,err := dbConn.Prepare("INSERT  INTO company (cid, uname, email, password, auth, register_date) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	tNow, _ := getCurrentTime()
	auth := "all"
	_, err = actIns.Exec(cid, uname, email, password, auth, tNow)
	if err != nil {
		return err
	}

	defer actIns.Close()
	return nil

}
