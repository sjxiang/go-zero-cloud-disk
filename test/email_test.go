package test

import (
	"testing"


	"net/smtp"
	"github.com/jordan-wright/email"
)


func TestSendEmail(t *testing.T) {
	e := email.NewEmail()

	e.From = "new@qq.com"
	e.To = []string{"old@qq.com"}
	e.Subject = "验证码发送测试"
	e.Text = []byte("123456")

	err := e.Send(
		"smtp.qq.com:25",
		smtp.PlainAuth(
			"",
			"new@qq.com",                   // 服务器邮箱账号
			"VERIFYCODE_QQEmailAuthCode",   // 授权码
			"smtp.qq.com",
		),
	)

	if err != nil {
		t.Errorf("发送邮件错误 %s", err.Error())
	}
}