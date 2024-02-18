package main

import (
	"noteapp/initializers"
	"noteapp/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.SyncDatabase()
}
func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routers.UserRouters(router)
	err := router.Run(":8000")
	if err != nil {
		panic(err.Error())
	}
}
