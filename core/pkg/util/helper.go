package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"

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
