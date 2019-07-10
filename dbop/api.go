package dbop

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func check(err error){
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
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
func userRegister(userName string, email string, password string) error {
	actIns,err := dbConn.Prepare("INSERT  INTO user_information (user_name, email, user_passwd) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = actIns.Exec(userName, email, password)
	if err != nil {
		return err
	}
	defer actIns.Close()
	return nil
}
