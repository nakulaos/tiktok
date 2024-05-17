package i18n

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"tiktok/common/errorcode"
)

//go:embed locale/*.json
var LocaleFS embed.FS

var langTagMap = map[string]language.Tag{
	"zh": language.Chinese,
	"en": language.English,
	"de": language.German,
	"jp": language.Japanese,
	"fr": language.French,
}

// Translator is a struct storing translating data.
type Translator struct {
	bundle       *i18n.Bundle
	localizer    map[language.Tag]*i18n.Localizer
	supportLangs []language.Tag
	conf         Conf
}

// NewBundle returns a bundle from FS.
func (l *Translator) NewBundle(files embed.FS) {
	filenames := GetFilenames(files)
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, filename := range filenames {
		_, err := bundle.LoadMessageFileFS(files, fmt.Sprintf("%s", filename))
		lang := strings.Split(filename, "/")[1]
		lang = strings.Split(lang, ".")[0]

		l.supportLangs = append(l.supportLangs, parseTags(lang)[0])
		l.conf.Language = append(l.conf.Language, lang)

		logx.Must(err)
	}
	l.bundle = bundle
}

// GetFilenames returns a slice of filenames from the provided FS.
func GetFilenames(file embed.FS) []string {
	var filenames []string
	fs.WalkDir(file, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			filenames = append(filenames, path)
		}
		return nil
	})
	return filenames
}

// NewBundleFromFile returns a bundle from a directory which contains resource files.
func (l *Translator) NewBundleFromFile(conf Conf) {
	bundle := i18n.NewBundle(language.Chinese)
	filePath, err := filepath.Abs(conf.Dir)
	logx.Must(err)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, lang := range conf.Language {
		_, err = bundle.LoadMessageFile(filepath.Join(filePath, fmt.Sprintf("%s.json", lang)))

		if lang != "zh" && lang != "en" {
			l.supportLangs = append(l.supportLangs, parseTags(lang)[0])
		}

		logx.Must(err)
	}

	l.bundle = bundle
	l.conf = conf
}

// NewTranslator sets localize for translator.
func (l *Translator) NewTranslator() {
	l.supportLangs = append(l.supportLangs, language.Chinese)
	l.supportLangs = append(l.supportLangs, language.English)
	l.localizer = make(map[language.Tag]*i18n.Localizer)
	for _, lang := range l.conf.Language {
		l.localizer[parseTags(lang)[0]] = i18n.NewLocalizer(l.bundle, parseTags(lang)[0].String())
	}
}

// Trans used to translate any resource string.
func (l *Translator) Trans(ctx context.Context, msgId string) string {
	message, err := l.MatchLocalizer(ctx.Value("lang").(string)).LocalizeMessage(&i18n.Message{ID: msgId})
	if err != nil {
		return msgId
	}

	if message == "" {
		return msgId
	}

	return message
}

// Trans used to translate any resource string.
func (l *Translator) TransWithTemplateData(ctx context.Context, msgId string, data map[string]interface{}) string {
	message, err := l.MatchLocalizer(ctx.Value("lang").(string)).Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: msgId,
		},
		TemplateData: data,
	})
	if err != nil {
		return msgId
	}
	if message == "" {
		return msgId
	}
	return message
}

// TransError translates the error message
func (l *Translator) TransError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	lang := ctx.Value("lang").(string)
	if codeErr, ok := err.(errorcode.ErrorCode); ok {
		details := codeErr.Details()
		if len(details) == 0 {
			message, e := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: codeErr.Error()})
			if e != nil || message == "" {
				message = codeErr.Error()
			}
			return errorcode.New(codeErr.Code(), message)
		} else {
			for _, detail := range details {
				switch detailInfo := detail.(type) {
				case *errdetails.ErrorInfo:
					//var templateData = make(map[string]interface{})
					//for key, val := range detailInfo.Metadata {
					//	templateData[key] = val
					//}
					message, e := l.MatchLocalizer(lang).Localize(&i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID: codeErr.Error(),
						},
						TemplateData: detailInfo.Metadata,
					})
					if e != nil || message == "" {
						message = codeErr.Error()
					}
					return errorcode.New(codeErr.Code(), message)
				case *errdetails.BadRequest_FieldViolation:
					// 暂时只返回第一个错误
					message, e := l.MatchLocalizer(lang).Localize(&i18n.LocalizeConfig{
						DefaultMessage: &i18n.Message{
							ID: codeErr.Error(),
						},
						TemplateData: map[string]interface{}{
							"field":       detailInfo.Field,
							"description": detailInfo.Description,
						},
					})
					if e != nil || message == "" {
						message = codeErr.Error()
					}
					return errorcode.New(codeErr.Code(), message)
				}

			}

			message, e := l.MatchLocalizer(lang).LocalizeMessage(&i18n.Message{ID: codeErr.Error()})
			if e != nil || message == "" {
				message = codeErr.Error()
			}
			return errorcode.New(codeErr.Code(), message)

		}

	} else {
		return errorcode.New(http.StatusInternalServerError, "failed to translate error message")
	}
	//else if apiErr, ok := err.(*errorx.ApiError); ok {
	//	message, e := l.MatchLocalizer(resource).LocalizeMessage(&resource.Message{ID: apiErr.Error()})
	//	if e != nil {
	//		message = apiErr.Error()
	//	}
	//	return errorx.NewApiError(apiErr.Code, message)
	//}
}

// MatchLocalizer used to matcher the localizer in map
func (l *Translator) MatchLocalizer(lang string) *i18n.Localizer {
	tags := parseTags(lang)
	for _, v := range tags {
		if val, ok := l.localizer[v]; ok {
			return val
		}
	}

	return l.localizer[language.Chinese]
}

// NewTranslator returns a translator by FS.
func NewTranslator(file embed.FS) *Translator {
	trans := &Translator{}
	trans.NewBundle(file)
	trans.NewTranslator()
	return trans
}

// NewTranslatorFromFile returns a translator by FS.
func NewTranslatorFromFile(conf Conf) *Translator {
	trans := &Translator{}
	trans.NewBundleFromFile(conf)
	trans.NewTranslator()
	return trans
}

func parseTags(lang string) []language.Tag {
	tags, _, err := language.ParseAcceptLanguage(lang)
	if err != nil {
		logx.Errorw("parse accept-language failed", logx.Field("detail", err))
		return []language.Tag{language.Chinese}
	}

	return tags
}
