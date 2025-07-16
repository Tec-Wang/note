package logic

import (
	"context"
	"fmt"
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
	// Validation: Check if files exist
	if len(l.Files) == 0 {
		return nil, fmt.Errorf("no file provided")
	}
	
	// Validation: Check if filename is provided
	if l.FileName == "" {
		return nil, fmt.Errorf("filename is required")
	}
	
	reader := l.Files[0]

	// Fix: Use the actual filename instead of hardcoded one
	err = l.svcCtx.FileStorage.UploadFile(l.FileName, reader)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}
	
	// Return success response
	return &types.NoteUploadResponse{
		Base: types.Base{
			Code: 200,
			Msg:  "success",
		},
		Data: types.NoteUploadResponseData{
			ID: 1, // This should be the actual ID from database
		},
	}, nil
}
