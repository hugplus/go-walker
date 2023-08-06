package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	OK    = 200
	ERROR = 500
)

type resp struct {
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
	c.JSON(http.StatusOK, resp{
		ReqId: c.GetString("reqId"),
		Code:  OK,
		Msg:   "ok",
		Data:  data,
	})
}

func Fail(c *gin.Context, code int, msg string, data ...any) {
	c.JSON(http.StatusOK, resp{
		ReqId: c.GetString("reqId"),
		Code:  ERROR,
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
	c.JSON(http.StatusOK, resp{
		ReqId: c.GetString("reqId"),
		Code:  OK,
		Msg:   "ok",
		Data:  p,
	})
}
