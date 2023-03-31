package controllers

import (
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			order := v1.Group("orders")
			{
				order.POST("", CreateOrder)
				order.GET("/:id", GetOrder)
				order.GET("", GetOrders)
				order.PUT("/:id", UpdateOrder)
			}
		}
	}
}
