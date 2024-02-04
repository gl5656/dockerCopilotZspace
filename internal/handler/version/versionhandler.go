package version

import (
	"net/http"

	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/logic/version"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/svc"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VersionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := version.NewVersionLogic(r.Context(), svcCtx)
		resp, err := l.Version(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}