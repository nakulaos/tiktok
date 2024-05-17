package resource

import (
	"embed"
)

//go:embed locale/*.json
var LocaleFS embed.FS
