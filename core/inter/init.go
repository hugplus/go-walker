package inter

import (
	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/middleware"
	"github.com/hugplus/go-walker/config"
)

func Init(r *gin.Engine, cfg *config.AppCfg) {
	middleware.InitMiddleware(r, cfg)
}
