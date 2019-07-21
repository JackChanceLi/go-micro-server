package dbop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(114.116.180.115:33306)/user_information")
}


