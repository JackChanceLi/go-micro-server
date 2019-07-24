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
	//t.Run("register", testUserRegister)
	//t.Run("login", testUserLogin)
	//t.Run("time",testGetCurrentTime)
	t.Run("latest_register",testNewUserRegister)
}

func testUserRegister(t *testing.T) {  // testing old register function
	err := userRegister("zheng","1258@gmail.com","000000",1)
	if err != nil {
		t.Errorf("Error of register: %v", err)
	}
}

func testUserLogin(t *testing.T) {
	password,err := UserLogin("zheng")
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

func testNewUserRegister(t *testing.T) {  //testing latest register function
	err := UserRegister("ljc", "jackchance@163.com", "123456789")
	if err != nil {
		t.Errorf("Error of latest user register: %v",err)
	}
}

func TestUserLogin(t *testing.T) {

}