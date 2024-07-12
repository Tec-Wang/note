package utils

import (
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	// 设置SMTP服务器地址和端口
	smtpHost := "smtp.qq.com"
	smtpPort := "587"
	// 邮件发送者地址
	from := "wangzheng4j@qq.com"
	// 构建邮件正文
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	// 设置SMTP客户端
	auth := smtp.PlainAuth("", from, "blcrlzozdmpkihcc", smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}
