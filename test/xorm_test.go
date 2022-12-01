package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/sjxiang/go-zero-cloud-disk/models"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func TestXORM(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(172.20.0.1:3306)/cloud-disk?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)  // 映射到 struct
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(data) // 转换成 bytes 数组
	if err != nil {
		t.Fatal(err)
	}

	dst := new(bytes.Buffer)  
	err = json.Indent(dst, b, "", "  ")  // 转换成 bytes.Buffer，前缀 间隔 
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dst.String())  // 转换成字符串，展示
}
