package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func sql_connect() {
	fmt.Println("helloworld")
	db, err := sql.Open("mysql", "root:123456@tcp(114.116.180.115:33306)/user_information")
	check(err)

	//插入数据
	//stmt, err := db.Prepare("INSERT  INTO user_information (user_name, email, user_passwd) VALUES (?, ?, ?)")
	//check(err)
	//res, err := stmt.Exec("zht","1023@gmail.com","123456")
	//check(err)
	//id, err := res.LastInsertId()
	//check(err)
	//fmt.Println(id)

	//打印数据
	rows, err := db.Query("SELECT * FROM user_information")
	check(err)
	i := 1
	for rows.Next() {
		var uid string
		var username string
		var passwd string
		var company string
		var date string
		var email string
		err = rows.Scan(&uid, &username, &passwd, &company, &date, &email)
		check(err)
		fmt.Println("NO:%d",i)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(passwd)
		fmt.Println(company)
		fmt.Println(date)
		fmt.Println(email)
	}
	rows.Close()

}

func check(err error){
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}


