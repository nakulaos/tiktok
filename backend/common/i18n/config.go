package i18n

// Conf is the configuration structure for resource
type Conf struct {
	Dir      string   `json:",env=I18N_DIR"`
	Language []string `json:",env=I18N_LANGUAGE"`
}
