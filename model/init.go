package model

import (
	"log"

	"xorm.io/xorm"
	_ "github.com/go-sql-driver/mysql"
)


// var Engine = Init()

func Init(dsn string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Printf("xorm New Engine Error: %v", err)
		return nil
	}

	return engine
}