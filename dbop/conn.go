package dbop

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	dbConn *sql.DB
	err error
)

func init(){
	dbConn, err = sql.Open("mysql", "root:123456@tcp(114.116.180.115:33306)/user_information")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
