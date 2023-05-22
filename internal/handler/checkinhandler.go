package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"hotel/internal/logic"
	"hotel/internal/svc"
	"hotel/internal/types"
	"template/response"
)

func CheckInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckInReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCheckInLogic(r.Context(), svcCtx)
		resp, err := l.CheckIn(&req)
		response.Response(w, resp, err)
	}
}
