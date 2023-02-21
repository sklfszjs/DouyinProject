package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnPool struct {
	DB chan *gorm.DB
}

var DBPool *DBConnPool
var capi = 50

func init() {
	DBPool = &DBConnPool{}
	DBPool.DB = make(chan *gorm.DB, capi)
	var err error
	var db *gorm.DB
	for i := 0; i < capi; i++ {
		db, err = gorm.Open(
			mysql.Open(fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				GetConfigs().DBUsrName, GetConfigs().DBPassword, GetConfigs().DBName)),
			&gorm.Config{})
		if err != nil {
			fmt.Println("数据库链接错误", err)
		}
		DBPool.DB <- db
	}

}

func GetDBConnPool() *DBConnPool {
	return DBPool
}

func (p *DBConnPool) GetConn() *gorm.DB {
	for {
		select {
		case conn := <-p.DB:
			fmt.Println("take out a conn, remain num:", len(p.DB))
			return conn
		}
	}
}

func (p *DBConnPool) ReturnConn(conn *gorm.DB) {
	p.DB <- conn
	fmt.Println("give back a conn, remain num:", len(p.DB))

}
