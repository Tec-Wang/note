package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wangzheng/brain/api/internal/config"
	"wangzheng/brain/model"
)

type ServiceContext struct {
	Config    config.Config
	NoteModel model.NoteModel
	FileModel model.FileModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	noteModel := model.NewNoteModel(sqlConn, c.Cache)
	fileModel := model.NewFileModel(sqlConn, c.Cache)
	return &ServiceContext{
		Config:    c,
		NoteModel: noteModel,
		FileModel: fileModel,
	}
}
