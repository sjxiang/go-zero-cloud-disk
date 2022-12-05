package test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)


var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "172.20.0.1:6379",
	Password: "",  // no password set
	DB: 0,         // use default DB
})

func TestSetValue(t *testing.T) {
	
	expiredTime := time.Second * 1024
	err := rdb.Set(ctx, "list", "1", expiredTime).Err()
	if err != nil {
		t.Error(err)
	}
}


func TestGetValue(t *testing.T) {

	val, err := rdb.Get(ctx, "list").Result()
	if err != nil {
		t.Error(err)
	}

	t.Log(val)
}