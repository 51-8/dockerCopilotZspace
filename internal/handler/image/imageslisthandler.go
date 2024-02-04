package image

import (
	"net/http"

	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/logic/image"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ImagesListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := image.NewImagesListLogic(r.Context(), svcCtx)
		resp, err := l.ImagesList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
