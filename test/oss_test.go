package test

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestFileUploadByFilepath(t *testing.T) {

	// 访问域名，存储桶名称 + 用户所在 region 
    u, _ := url.Parse("")
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            // 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
            SecretID: "",
            // 表示用户的 SecretKey
            SecretKey: "",
        },
    })

    key := "cloud-disk/jisoo.jpg"  // oss

    _, _, err := client.Object.Upload(
        context.Background(), key, "./img/jisoo.jpg", nil,  // local server
    )
    if err != nil {
        t.Error(err)
    }
}

func TestFileUploadByReader(t *testing.T) {

	// 访问域名，存储桶名称 + 用户所在 region 
    u, _ := url.Parse("")
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            // 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
            SecretID: "",
            // 表示用户的 SecretKey
            SecretKey: "",
        },
    })

    key := "cloud-disk/jisoo2.jpg"  // oss

	fd, err := os.ReadFile("./img/jisoo.jpg")
	if err != nil {
        t.Error(err)
    }
	
	_, err = client.Object.Put(
        context.Background(), key, bytes.NewBuffer(fd), nil,  // local server
    )
    if err != nil {
        t.Error(err)
    }
}