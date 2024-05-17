package {{.PkgName}}

import (
	"net/http"
	"tiktok/common/base"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := base.Parse(r, &req); err != nil {
		    err = svcCtx.Trans.TransError(r.Context(), err)
			base.HttpResult(r,w,nil, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		err = svcCtx.Trans.TransError(r.Context(),err)
		base.HttpResult(r, w, {{if .HasResp}}resp{{else}}nil{{end}}, err)
	}
}
