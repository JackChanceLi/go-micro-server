package dbop

import (
	_ "github.com/go-sql-driver/mysql"
	"go-micro-server/utils"
)

func InsertLiveRoom(cid string, name string, kind int, size int, start_time string, end_time string, status string, permission string, create_time string) error {
	stmtIns,err := dbConn.Prepare("INSERT INTO live_room (lid, cid, name, kind, size, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, status, permission, create_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	lid, _ := utils.NewUUID()
	push_url := "www.baidu.com"
	pull_hls_url := "www.baidu.com"
	pull_rtmp_url := "www.baidu.com"
	pull_http_flv_url := "www.baidu.com"
	display_url := "www.baidu.com"

	_,err = stmtIns.Exec(lid, cid, name, kind, size, start_time, end_time, push_url, pull_hls_url, pull_rtmp_url, pull_http_flv_url, display_url, status, permission, create_time)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return  nil
}

func DeleteLiveRoom(cid string) {

}