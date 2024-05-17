package svc

import (
    "tiktok/common/i18n"
	{{.configImport}}
)

type ServiceContext struct {
	Config {{.config}}
	Trans  *i18n.Translator
	{{.middleware}}
}

func NewServiceContext(c {{.config}}) *ServiceContext {
	var trans *i18n.Translator
	if c.I18nConf.Dir != "" {
		trans = i18n.NewTranslatorFromFile(c.I18nConf)
	}
	return &ServiceContext{
		Config: c,
		Trans: trans,
		{{.middlewareAssignment}}
	}
}
