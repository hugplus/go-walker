package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/consts"
	"github.com/hugplus/go-walker/common/errs"
	"github.com/hugplus/go-walker/common/errs/codes"
)

type Resp struct {
	ReqId string `json:"reqId"`          //`json:"请求id"`
	Code  int    `json:"code"`           //返回码
	Msg   string `json:"msg,omitempty"`  //消息
	Data  any    `json:"data,omitempty"` //数据
}

type PageResp struct {
	List  []any `json:"list"`  //数据列表
	Total int   `json:"total"` //总条数
	Size  int   `json:"size"`  //分页大小
	Page  int   `json:"list"`  //当前第几页
}

//type RespFunc func()

func Ok(c *gin.Context, data any) {
	c.AbortWithStatusJSON(http.StatusOK, Resp{
		ReqId: c.GetString(consts.REQ_ID),
		Code:  codes.SUCCESS,
		Msg:   "ok",
		Data:  data,
	})
}

func Err(c *gin.Context, err errs.IError, msg string) {
	Fail(c, err.Code(), msg)
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
