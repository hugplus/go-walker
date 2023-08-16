package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/config"
)

func InitMiddleware(r *gin.Engine, cfg *config.AppCfg) {
	//.Use(Options)
	//.Use(Secure)
	if cfg.Cors.Enable {
		r.Use(CorsByRules(&cfg.Cors))
	}
	r.Use(CustomError)
	r.Use(ReqId)
	//r.Use(NoCache)

}
