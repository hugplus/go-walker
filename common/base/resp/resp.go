package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/consts"
)

const (
	OK_CODE    = 200
	ERROR_CODE = 500
	OK_MSG     = "ok"
)

type Resp struct {
	ReqId string `json:"reqId"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  any    `json:"data"`
}

type PageResp struct {
	List  []any
	Total int
	Size  int
	Page  int
}

type RespFunc func()

func Ok(c *gin.Context, data any) {
	c.AbortWithStatusJSON(http.StatusOK, Resp{
		ReqId: c.GetString(consts.REQ_ID),
		Code:  OK_CODE,
		Msg:   OK_MSG,
		Data:  data,
	})
}

func Fail(c *gin.Context, code int, msg string, data ...any) {
	c.AbortWithStatusJSON(http.StatusOK, Resp{
		ReqId: c.GetString(consts.REQ_ID),
		Code:  code,
		Msg:   msg,
		Data:  data,
	})
}

func Page(c *gin.Context, list []any, count int, page int, pageSize int) {
	p := PageResp{
		Page:  page,
		Total: count,
		Size:  pageSize,
		List:  list,
	}
	Ok(c, p)
}
