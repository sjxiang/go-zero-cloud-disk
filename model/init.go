package model

import (
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
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


func InitRedis(addr, password string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,  // no password set
		DB: db,              // use default DB
	})
}


func InitOSS() {
	
}