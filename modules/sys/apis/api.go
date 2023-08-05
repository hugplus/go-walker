package apis

import "github.com/gin-gonic/gin"

// type SysApi struct{

// }

// Ping Ping接口
// @Summary Ping接口
// @Description Ping接口
// @Tags Default
// @Success 200 {object} response.Response{data=string}} "{"code": 200, "data": [...]}"
// @Router /api/v1/ping [get]
// @Security Bearer
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": "pong",
	})
}
