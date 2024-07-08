package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wangzheng/brain/api/internal/logic"
	"wangzheng/brain/api/internal/svc"
	"wangzheng/brain/api/internal/types"
)

// 上传文件
func uploadNoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NoteUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadNoteLogic(r.Context(), svcCtx)
		resp, err := l.UploadNote(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
