package utils

import (
	"fmt"
	"net/smtp"
	"os"
	"regexp"

	"github.com/sirupsen/logrus"
)

func SendEmail(to, subject, body string) error {
	// 设置SMTP服务器地址和端口
	smtpHost := "smtp.qq.com"
	smtpPort := "587"
	// 邮件发送者地址
	from := os.Getenv("SMTP_FROM")
	if from == "" {
		from = "wangzheng4j@qq.com" // fallback, but should be set via env
	}
	
	// 从环境变量获取SMTP密码
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	if smtpPassword == "" {
		logrus.Error("SMTP_PASSWORD environment variable not set")
		return fmt.Errorf("SMTP credentials not configured")
	}
	
	// 构建邮件正文
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body
	// 设置SMTP客户端
	auth := smtp.PlainAuth("", from, smtpPassword, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}

func CheckEmail(email string) bool {
	// 正则表达式匹配邮箱格式
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// 修复：使用正则表达式包中的MatchString函数
	res, err := regexp.MatchString(pattern, email)
	if err != nil {
		logrus.Error(err)
		return false
	}
	return res
}
