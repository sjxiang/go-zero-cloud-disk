package util

import (
	"bytes"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"

	uuid "github.com/satori/go.uuid"

	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/mail"
)

func MD5(plainText string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(plainText)))
}

func VerifyCodeSend(email, code string) error {
	if ok := mail.NewMailer().Send(mail.Email{
		From: os.Getenv("VERIFYCODE_FROM"),
		To: []string{email},
		Subject: "cloud-disk 验证码",
		Text: []byte(fmt.Sprintf("您的 Email 验证码：%s", code)),
	}); !ok {
		return errors.New("email 发送失败")
	}

	return nil	
}


// RandStringRunes 生成长度为 length 随机数字字符串
func RandStringRunes(length int) string {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")  

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	
	return string(b)
}


func RandVerifyCode() string {
	codeLength := 6
	return RandStringRunes(codeLength)
}

func GenUUID() string {
	return uuid.NewV4().String()
}


// 文件上传到腾讯云
func OSSUpload(r *http.Request) (string, error) {
	
	// 访问域名，存储桶名称 + 用户所在 region 
    u, _ := url.Parse(os.Getenv("BUCKETURL"))
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
           	SecretID: os.Getenv("SECRETID"),
            SecretKey: os.Getenv("SECRETKEY"),
        },
    })

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", err
	}

    key := "cloud-disk/" + GenUUID() + path.Ext(fileHeader.Filename)  // oss


    _, err = client.Object.Put(
		context.Background(), key, file, nil, 
	)
	if err != nil {
		return "", err
	}

	return os.Getenv("BUCKETURL") + "/" + key, nil 
}

func OSSInitPart(ext string) (string, string, error) {
	u, _ := url.Parse(os.Getenv("BUCKETURL"))
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
           	SecretID: os.Getenv("SECRETID"),
            SecretKey: os.Getenv("SECRETKEY"),
        },
    })
	key := "cloud-disk/" + GenUUID() + ext  // oss
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}
	UploadID := v.UploadID
	
	return key, UploadID, nil
}


// 分片上传
func OSSPartUpload(r *http.Request) (string, error) {

	u, _ := url.Parse(os.Getenv("BUCKETURL"))
    b := &cos.BaseURL{BucketURL: u}
    client := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
           	SecretID: os.Getenv("SECRETID"),
            SecretKey: os.Getenv("SECRETKEY"),
        },
    })

	key := r.PostForm.Get("key")
	uploadID := r.PostForm.Get("uploadId")
	partNumber, _ := strconv.Atoi(r.PostForm.Get("partNumber"))

	f, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)  // Content-Length HTTP header 

	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, uploadID, partNumber, bytes.NewReader(buf.Bytes()), nil,  // 需要 io.Reader
	)
	if err != nil {
		return "", err
	}

	PartETag := resp.Header.Get("ETag")

	return strings.Trim(PartETag, "\""), nil
}