package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wangzheng/brain/api/internal/logic"
	"wangzheng/brain/api/internal/svc"
)

const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)

// 上传文件
func uploadNoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadNoteLogic(r.Context(), svcCtx)

		file, _, err := r.FormFile("fileContent")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l.Files = append(l.Files, file)
		l.FileName = r.FormValue("filename")

		resp, err := l.UploadNote(nil)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
