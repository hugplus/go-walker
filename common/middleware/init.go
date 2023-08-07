package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	r.Use(Options).Use(NoCache).Use(Secure).Use(ReqId())
}
