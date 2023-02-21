package utils

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestPool(t *testing.T) {
	connpool := GetDBConnPool()
	var dblist [10]*gorm.DB
	for i := 0; i < 8; i++ {
		dblist[i] = connpool.GetConn()
		time.Sleep(time.Second * 1)
	}
	for i := 0; i < 5; i++ {
		connpool.ReturnConn(dblist[i])
		time.Sleep(time.Second * 1)
	}
}
