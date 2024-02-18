package routers

import (
	"noteapp/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/register", controllers.RegitserUser())
	incomingRoutes.POST("/login", controllers.Login())

}
