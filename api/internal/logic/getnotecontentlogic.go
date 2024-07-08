package logic

import (
	"context"

	"wangzheng/brain/api/internal/svc"
	"wangzheng/brain/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoteContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文件内容
func NewGetNoteContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoteContentLogic {
	return &GetNoteContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoteContentLogic) GetNoteContent(req *types.Id) (resp *types.NoteContentResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
