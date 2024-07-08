package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wangzheng/brain/api/internal/logic"
	"wangzheng/brain/api/internal/svc"
	"wangzheng/brain/api/internal/types"
)

// 获取文件内容
func getNoteContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Id
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetNoteContentLogic(r.Context(), svcCtx)
		resp, err := l.GetNoteContent(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
