package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/consts"
	"github.com/hugplus/go-walker/common/errs"
	"github.com/hugplus/go-walker/common/errs/codes"
)

type Resp struct {
	ReqId string `json:"reqId"`
	Code  int    `json:"code"`
	Msg   string `json:"msg,omitempty"`
	Data  any    `json:"data,omitempty"`
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
		Code:  codes.SUCCESS,
		Msg:   codes.MSG_OK,
		Data:  data,
	})
}

func Err(c *gin.Context, err errs.IError) {
	Fail(c, err.Code(), err.Msg())
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
