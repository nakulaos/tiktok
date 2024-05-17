/**
 ******************************************************************************
 * @file           : test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/4/8
 ******************************************************************************
 */

package main

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("active.zh.toml")

	localizer := i18n.NewLocalizer(bundle, "zh")

	helloMessage, _ := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "hello",
			Other: "Hello",
		},
	})

	fmt.Println(helloMessage) // 输出: 你好
}
