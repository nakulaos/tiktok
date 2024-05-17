package config

import (
    ii18n "tiktok/common/i18n"
    {{.authImport}}
)

type Config struct {
	rest.RestConf
	{{.auth}}
	{{.jwtTrans}}
	I18nConf ii18n.Conf
}
