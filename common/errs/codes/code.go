package codes

import (
	"github.com/hugplus/go-walker/core"
	"golang.org/x/text/language"
)

const (
	SUCCESS            = 200
	FAILURE            = 500
	AuthorizationError = 403
	NotFound           = 404
	NotLogin           = 401
	InvalidParameter   = 10000
	UserDoesNotExist   = 10001
	ServerError        = 10101
	TooManyRequests    = 10102
)

func GetLangMsgByCode(acceptLanguate string, code int) string {
	if core.Cfg.Server.I18n {
		tags, _, _ := language.ParseAcceptLanguage(acceptLanguate)
		if len(tags) > 0 {
			return GetMsg(code, tags[0].String())
		}
	}
	return GetMsg(code, core.Cfg.Server.GetLang())
}

const (
	LANG_ZH_CN = "zh-CN"
	LANG_ZH    = "zh"
	LANG_EN    = "en"
)

func GetMsg(code int, lang string) (str string) {
	var ok bool
	switch lang {
	case LANG_ZH_CN, LANG_ZH:
		str, ok = zhCNText[code]
	case LANG_EN:
		str, ok = enUSText[code]
	default:
		str, ok = zhCNText[code]
	}
	if !ok {
		return "unknown error"
	}
	return
}
