package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hugplus/go-walker/common/consts"
)

type BaseApi struct {
}

func (e *BaseApi) GetReqId(c *gin.Context) string {
	reqId := c.GetString(consts.REQ_ID)
	if reqId == "" {
		reqId = uuid.NewString()
		c.Set(consts.REQ_ID, reqId)
	}
	return reqId
}

func (e *BaseApi) Bind(c *gin.Context, o *any) {
	c.ShouldBind(o)
	
}
