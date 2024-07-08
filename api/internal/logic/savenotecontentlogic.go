package logic

import (
	"context"

	"wangzheng/brain/api/internal/svc"
	"wangzheng/brain/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveNoteContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 保存文件内容
func NewSaveNoteContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveNoteContentLogic {
	return &SaveNoteContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveNoteContentLogic) SaveNoteContent(req *types.NoteUploadRequest) (resp *types.Base, err error) {
	// todo: add your logic here and delete this line

	return
}
