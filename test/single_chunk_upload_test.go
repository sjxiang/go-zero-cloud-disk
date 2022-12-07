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

// 分片上传初始化
func TestInitPartUpload(t *testing.T) {

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

    key := "cloud-disk/chunk-jisoo.jpg"  // oss

	// 可选opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Error(err)
	}
	UploadID := v.UploadID  //  提前占坑的房卡
	t.Log(UploadID)
}

// 分片上传
func TestPartUpload(t *testing.T) {
	u, _ := url.Parse("om")
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            // 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: "",
            // 表示用户的 SecretKey
            SecretKey: "",
		},
    })

    key := "cloud-disk/chunk-jisoo.jpg"  // oss
	UploadID := ""
	
	
	bs, err := os.ReadFile("./img/0_chunk")  // md5 "413d439e6306e432"
	if err != nil {
		t.Error(err)
	}

	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, bytes.NewReader(bs), nil,  // 循环
	)
	if err != nil {
		t.Error(err)
	}

	PartETag := resp.Header.Get("ETag")
	t.Log(PartETag)
}


// 分片上传完成
func TestPartUploadComplete(t *testing.T) {
	
	u, _ := url.Parse("htmyqcloud.com")
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            // 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
            SecretID: "",
            // 表示用户的 SecretKey
            SecretKey: "",
        },
    })

    key := "cloud-disk/chunk-jisoo.jpg"  // oss
	UploadID := "309c"
	


	opt := &cos.CompleteMultipartUploadOptions{}
    opt.Parts = append(opt.Parts, cos.Object{
        PartNumber: 1, ETag: "2"},
    )
    _, _, err := client.Object.CompleteMultipartUpload(
        context.Background(), key, UploadID, opt,
    )
	if err != nil {
		t.Error(err)
	}
}