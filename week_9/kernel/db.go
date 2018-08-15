package kernel

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DB struct {
	Conn *gorm.DB
}

func (d *DB) CreateConn() {
	user := "root"
	pw := "root"
	ip := "127.0.0.1:3306"
	db := "test"
	c, err := gorm.Open("mysql", user+":"+pw+"@tcp("+ip+")/"+db)
	if err != nil {
		fmt.Printf("Connect to Mysql error: %+v\n", err)
	}

	d.Conn = c
}
