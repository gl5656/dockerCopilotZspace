package handler

import (
	"github.com/flosch/pongo2"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"

	"github.com/onlyLTY/oneKeyUpdate/v2/internal/logic"
	"github.com/onlyLTY/oneKeyUpdate/v2/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginIndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLoginIndexLogic(r.Context(), svcCtx)
		err := l.LoginIndex()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			t, _ := svcCtx.Template.FromFile("templates/login/login.html")
			//if err != nil {
			//	logx.Error(err)
			//}
			execute, err := t.ExecuteBytes(pongo2.Context{"current_year": time.Now()})
			if err != nil {
				logx.Error(err)
			}
			_, err = w.Write(execute)
			if err != nil {
				logx.Error(err)
			}
			//cookies := &http.Cookie{
			//	Name:   "device_verified",
			//	Value:  "true",
			//	MaxAge: 365 * 24 * 60 * 60,
			//}
			//w.Header().Set("Set-Cookie", cookies.String())
			httpx.Ok(w)
		}
	}
}