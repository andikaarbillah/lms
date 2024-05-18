package config

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

const (
	host     = "localhost"
	port     = 3307
	dbname   = "lms"
	username = "root"
	password = ""
)

var dsn = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)

const (
    colorGreen = "\033[1;32m"
    colorRed   = "\033[1;31m"
    colorReset = "\033[0m"
)

func DB() *gorm.DB {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        fmt.Printf("\n----------------\n|%sNOT CONNECTED!%s|\n----------------\n", colorRed, colorReset)
        panic(err)
    }
   		fmt.Printf("\n----------------\n|%sIS-->CONNECTED%s|\n----------------\n", colorGreen, colorReset)

    return db
}