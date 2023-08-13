package base

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/errs/codes"
)

func FmtReqId(reqId string) string {
	return fmt.Sprintf("REQID:%s", reqId)
}

func GetAcceptLanguage(c *gin.Context) string {
	return c.GetHeader("Accept-Language")
}

func GetMsgByCode(c *gin.Context, code int) string {
	return codes.GetLangMsgByCode(GetAcceptLanguage(c), code)
}
