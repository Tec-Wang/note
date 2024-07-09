package logic

import (
	"context"
	"io"

	"wangzheng/brain/api/internal/svc"
	"wangzheng/brain/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadNoteLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	Files    []io.Reader
	FileName string
}

// 上传文件
func NewUploadNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadNoteLogic {
	return &UploadNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadNoteLogic) UploadNote(req *types.NoteUploadRequest) (resp *types.NoteUploadResponse, err error) {
	reader := l.Files[0]

	err = l.svcCtx.FileStorage.UploadFile("测试.txt", reader)
	return nil, err
}
