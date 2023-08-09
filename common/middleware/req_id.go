package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/utils"
)

func ReqId(c *gin.Context) {
	utils.GetReqId(c)
	c.Next()
}

// func ReqId() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		reqId := c.GetString(consts.REQ_ID)
// 		if reqId == "" {
// 			reqId = uuid.NewString()
// 			c.Set(consts.REQ_ID, reqId)
// 		}
// 		c.Next()
// 	}
// }
