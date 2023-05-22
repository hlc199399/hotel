package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"hotel/internal/logic"
	"hotel/internal/svc"
	"hotel/internal/types"
	"template/response"
)

func CompleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CompleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCompleteLogic(r.Context(), svcCtx)
		resp, err := l.Complete(&req)
		response.Response(w, resp, err)
	}
}
