package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sjxiang/go-zero-cloud-disk/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func TestGORM(t *testing.T) {
	
	dsn := "root:123456@tcp(172.20.0.1:3306)/cloud-disk?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.UserBasic, 0)
	err = db.First(&data, 1).Error // 根据整型主键查找
	
	if err != nil {

		// First vs. Find
		// 单个结构体 与 结构体切片，nil
		//
		// 查询结果 
		// 1. 0, gorm.ErrNotFound
		// 2. 0, err.Error()
		// 3. len(), nil

		t.Fatal(err.Error())
	}
	
	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dst.String())
}
