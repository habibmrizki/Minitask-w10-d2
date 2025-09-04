package routers

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/habibmrizki/gin/internal/handler"
)

func initPingRouter(router *gin.Engine) {
	pingRouter := router.Group("/ping")
	ph := handlers.NewPingHandler()

	pingRouter.GET("", ph.GetPing)
	pingRouter.GET("/:id/:param2", ph.GetPingWithParam)
	pingRouter.POST("", ph.PostPing)
}
