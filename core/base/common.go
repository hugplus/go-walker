package base

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/codes"
	"github.com/hugplus/go-walker/core"
	"golang.org/x/text/language"
)

func FmtReqId(reqId string) string {
	return fmt.Sprintf("REQID:%s", reqId)
}

func GetAcceptLanguage(c *gin.Context) string {
	return c.GetHeader("Accept-Language")
}

func GetMsgByCode(c *gin.Context, code int) string {
	if core.Cfg.Server.I18n {
		acceptLanguate := GetAcceptLanguage(c)
		tags, _, _ := language.ParseAcceptLanguage(acceptLanguate)
		if len(tags) > 0 {
			return codes.GetMsg(code, tags[0].String())
		}
	}
	return codes.GetMsg(code, core.Cfg.Server.GetLang())
}
