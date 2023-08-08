package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hugplus/go-walker/common/consts"
)

func ReqId(c *gin.Context) {
	reqId := c.GetString(consts.REQ_ID)
	if reqId == "" {
		reqId = uuid.NewString()
		c.Set(consts.REQ_ID, reqId)
	}
	fmt.Println("ReqId Bef")
	c.Next()
	fmt.Println("ReqId AFT")
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
