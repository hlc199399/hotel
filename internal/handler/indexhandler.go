package handler

import (
	"net/http"

	"hotel/internal/logic"
	"hotel/internal/svc"
	"template/response"
)

func IndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewIndexLogic(r.Context(), svcCtx)
		resp, err := l.Index()
		response.Response(w, resp, err)
	}
}
