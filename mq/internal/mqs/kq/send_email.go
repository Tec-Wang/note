package kq

import (
	"context"
	"encoding/json"
	"wangzheng/brain/common/kq"
	"wangzheng/brain/common/utils"
	"wangzheng/brain/mq/internal/svc"
)

type Email struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailSender(ctx context.Context, svcCtx *svc.ServiceContext) *Email {
	return &Email{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Email) Consume(_, val string) error {
	// 序列化消息
	msg := &kq.EmailMsg{}
	if err := json.Unmarshal([]byte(val), msg); err != nil {
		return err
	}

	return utils.SendEmail(msg.To, msg.Subject, msg.Body)
}
