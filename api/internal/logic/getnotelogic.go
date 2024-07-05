package logic

import (
	"context"

	"wangzheng/brain/api/internal/svc"
	"wangzheng/brain/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoteLogic {
	return &GetNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoteLogic) GetNote(req *types.ID) (resp *types.NoteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
