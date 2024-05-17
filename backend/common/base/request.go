package base

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"tiktok/common/errorcode"
	"tiktok/common/xvalidate"
)

var (
	validatorVal = xvalidate.NewValidator()
)

func Parse(r *http.Request, v any) error {

	if err := httpx.Parse(r, v); err != nil {
		return err
	}

	if errMsg := validatorVal.Validate(v); errMsg != "" {
		return errorcode.NewCodeInvalidArgumentError(errMsg)
	}

	return nil

}
