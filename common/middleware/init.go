package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	r.Use(ReqId)
	//r.Use(NoCache).Use(ReqId)
	//.Use(Options)

	//.Use(Secure)

}
