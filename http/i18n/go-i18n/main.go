package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func Test01(messageId string, lang int) string {
	// 创建 Bundle
	bundle := i18n.NewBundle(language.English)

	// 注册解析器
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	//加载翻译文件
	//translations := []string{
	//	"messages.en.toml",    // 英语翻译文件
	//	"messages.zh-CN.toml", // 简体中文翻译文件
	//}
	//for _, file := range translations {
	//	bundle.MustLoadMessageFile(file)
	//}

	// 创建 Localize
	localize := new(i18n.Localizer)
	if lang == 1 {
		bundle.MustLoadMessageFile("messages.zh-CN.toml")
		localize = i18n.NewLocalizer(bundle, "zh-CN")
	} else {
		bundle.MustLoadMessageFile("messages.en.toml")
		localize = i18n.NewLocalizer(bundle, "en")
	}

	// 执行翻译
	translation := localize.MustLocalize(&i18n.LocalizeConfig{
		MessageID: messageId,
	})

	return translation
}

func Test02(messageId, name string, count, lang int) string {
	// 创建 Bundle
	bundle := i18n.NewBundle(language.English)

	// 注册解析器
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// 加载翻译文件
	translations := []string{
		"messages.en.toml",    // 英语翻译文件
		"messages.zh-CN.toml", // 简体中文翻译文件
	}
	for _, file := range translations {
		bundle.MustLoadMessageFile(file)
	}

	// 创建 Localizer
	localize := new(i18n.Localizer)
	if lang == 1 {
		localize = i18n.NewLocalizer(bundle, "zh-CN")
	} else {
		localize = i18n.NewLocalizer(bundle, "en")
	}

	// 执行翻译
	translation := localize.MustLocalize(&i18n.LocalizeConfig{
		MessageID: messageId,
		TemplateData: map[string]interface{}{
			"Name":  name,
			"Count": count,
		},
		PluralCount: count,
	})

	return translation
}
func main() {
	message01 := Test01("Tips", 1)
	message02 := Test02("Emails", "LiMing", 2, 1)
	fmt.Println(message01)
	fmt.Println(message02)
}
