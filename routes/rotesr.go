package routes

import (
	"ginblog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//InitRouter 初始路由器
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	r.Run(utils.HTTPPort)
}
