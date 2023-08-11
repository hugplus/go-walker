package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	//.Use(Options)
	//.Use(Secure)

	r.Use(CustomError)
	r.Use(ReqId)
	r.Use(NoCache)

}
