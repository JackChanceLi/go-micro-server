package dbop

import (
	"fmt"
	"testing"
)

func dbClear(){
	dbConn.Exec("truncate user_information")
	dbConn.Exec(("truncate user_identity_information"))
}

func TestMain(m *testing.M) {
	//dbClear()
	m.Run()
	//dbClear()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("register", testUserRegister)
	t.Run("login", testUserLogin)
	t.Run("time",testGetCurrentTime)
}

func testUserRegister(t *testing.T) {
	err := UserRegister("zheng","1258@gmail.com","000000",1)
	if err != nil {
		t.Errorf("Error of register: %v", err)
	}
}

func testUserLogin(t *testing.T) {
	password,err := userLogin("zheng")
	if password != "000000" {
		t.Errorf("Error of user login for wrong password:%s",password)
	}
	if err != nil {
		t.Errorf("Error of login: %v", err)
	}
}

func testGetCurrentTime(t *testing.T) {
	tNow, err:= getCurrentTime()
	fmt.Printf("Time now:%s\n", tNow)
	if err != nil {
		t.Errorf("Error of time format:%s", err)
	}
}