package dbop

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

func userLogin(loginName string,) (string, error) {
	actOut,err := dbConn.Prepare("SELECT user_passwd from user_information WHERE user_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "",nil
	}
	var password string
	err = actOut.QueryRow(loginName).Scan(&password)
	if err != nil && err != sql.ErrNoRows {
		return "",err
	}
	defer actOut.Close()
	return password,nil
}
func UserRegister(userName string, email string, password string, role int) error {
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
