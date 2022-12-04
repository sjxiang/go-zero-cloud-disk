package mail

import (
	"net/smtp"
	"os"
	"log"

	emailPKG "github.com/jordan-wright/email"
)

// SMTP 实现 mail.Driver interface
type SMTP struct{}

// Send 实现 mail.Driver interface 的 Send 方法
func (S *SMTP) Send(email Email) bool {
	
	e := emailPKG.NewEmail()

	e.From = email.From
	e.To = email.To
	e.Subject = email.Subject
	e.Text = email.Text

	err := e.Send(
		"smtp.qq.com:25",
		smtp.PlainAuth(
			"",
			os.Getenv("VERIFYCODE_FROM"),  // 服务器邮箱账号
			os.Getenv("VERIFYCODE_QQEmailAuthCode"),  // 授权码
			"smtp.qq.com",
		),
	)

	if err != nil {
		log.Printf("发送邮件错误 %s", err.Error())
		return false
	}

	return true
}