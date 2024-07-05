package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wangzheng/brain/api/internal/logic"
	"wangzheng/brain/api/internal/svc"
	"wangzheng/brain/api/internal/types"
)

func GetNoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ID
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetNoteLogic(r.Context(), svcCtx)
		resp, err := l.GetNote(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
