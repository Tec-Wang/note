package listen

import (
	"context"
	"wangzheng/brain/mq/internal/config"
	"wangzheng/brain/mq/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	kqMq "wangzheng/brain/mq/internal/mqs/kq"
)

// pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.EmailSender, kqMq.NewEmailSender(ctx, svcContext)),
		//.....
	}
}
