package main

import (
	"context"
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle = i18n.NewBundle(language.English)
)

func init() {
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("active.zh.toml")
}

type BaseError struct {
	ID             string
	DefaultMessage string
	TempData       map[string]interface{}
}

func (b BaseError) Error() string {
	return b.DefaultMessage
}

func (b BaseError) LocalizedID() string {
	return b.ID
}

func (b BaseError) TemplateData() map[string]interface{} {
	return b.TempData
}

type LocalizedError interface {
	error
	LocalizedID() string
	TemplateData() map[string]interface{}
}

type UserNotFoundErr struct {
	BaseError
}

func NewUserNotFoundErr(userID int) LocalizedError {
	msg := i18n.Message{
		ID:    "user_not_found",
		Other: "User not found {{.UserID}}",
	}
	e := UserNotFoundErr{}
	e.ID = msg.ID
	e.DefaultMessage = msg.Other
	e.TempData = map[string]interface{}{
		"UserID": userID,
	}
	return e
}

func GetLang(_ context.Context) string {
	return "zh"
}

func FormatErr(err error) string {
	lang := GetLang(context.Background())
	loc := i18n.NewLocalizer(bundle, lang)
	var i18nErr LocalizedError
	if errors.As(err, &i18nErr) {
		msg, _ := loc.Localize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    i18nErr.LocalizedID(),
				Other: i18nErr.Error(),
			},
			//MessageID: i18nErr.LocalizedID(),
			TemplateData: i18nErr.TemplateData(),
		})
		return msg
	}
	return err.Error()
}
