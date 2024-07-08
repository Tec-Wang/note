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
	entity, err := l.svcCtx.NoteModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return nil, err
	}

	one, err := l.svcCtx.FileModel.FindOne(l.ctx, entity.FileId.Int64)
	if err != nil {
		return nil, err
	}

	l.fileParseRpc.Parse(l.ctx, &types.FileID{ID: one.Id})

	creator, err := l.svcCtx.UserModel.FindOne(l.ctx, entity.CreatorId)
	if err != nil {
		return nil, err
	}

	c := types.Creator{
		ID:    entity.CreatorId,
		Name:  creator.Name,
		Email: creator.Email,
	}

	return &types.NoteResp{
		ID:      entity.Id,
		Title:   entity.Title.String,
		FileID:  entity.FileId.Int64,
		Creator: c,
	}, nil
}
