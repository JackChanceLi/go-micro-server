package utils

import (
	"fmt"
	"testing"
)

func TestNewUUID(t *testing.T) {
	uid, err := NewUUID()
	if err != nil {
		t.Errorf("Error of generate new id:%v",err)
	}
	fmt.Printf("New id is: %s\n",uid)
}
