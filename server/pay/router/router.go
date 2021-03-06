// @Author liuYong
// @Created at 2020-01-05
// @Modified at 2020-01-05
package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "pay/controller/api/v1"
	_ "pay/docs"
)

// @title 支付服务
// @version 1.0
// @description 负责处理与支付和退款相关的业务

// @contact.name LiuYong
// @contact.email ly_yong@qq.com

// @host localhost:8082
// @BasePath /pay/api/v1
// @query.collection.format multi
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := r.Group("/pay/api/v1")
	{
		okGroup := apiV1.Group("/ok")
		{
			okGroup.POST("/abb", v1.PayOKAbb)
		}
		wantGroup := apiV1.Group("/want")
		{
			wantGroup.POST("/abb", v1.WantPayAbb)
		}
		refundGroup := apiV1.Group("/refund")
		{
			refundGroup.POST("/abb", v1.RefundAbb)
		}
	}

	return r
}
