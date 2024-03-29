package Controllers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zxc10110/mvc_63050096_2565_1/View"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/createFeedback", CreateFeedback)
	r.GET("/getFeedback", GetFeedBack)
	r.POST("/updateFeedback", UpdateFeedback)
	r.POST("/adminUpdate", AdminUpdate)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
