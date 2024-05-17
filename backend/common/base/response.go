package base

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	xerr "github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	"google.golang.org/grpc/status"
	"net/http"
	"tiktok/common/errorcode"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
	} else {
		errCode := errorcode.ServerError.Code()
		errMsg := "Internal Server Error"

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(errorcode.ErrorCode); ok {
			errCode = e.Code()
			errMsg = e.Message()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok {
				errCode = uint32(gstatus.Code())
				errMsg = gstatus.Message()
			}

		}
		logx.WithContext(r.Context()).Errorf("API-ERR:%+v", err)
		xhttp.JsonBaseResponseCtx(r.Context(), w, xerr.New(int(errCode), errMsg))
	}
}
