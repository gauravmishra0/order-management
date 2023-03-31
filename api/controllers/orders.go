package controllers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"order-management-service/api/data/dao"
	"order-management-service/api/models"
	"order-management-service/api/utils"
)

func CreateOrder(c *gin.Context) {
	log := utils.GetLogCtx(c)
	var req models.CreateOrderRequest

	if err := c.BindJSON(&req); err != nil {
		log.Errorf("Invalid payload format %v", err)
		c.JSON(http.StatusBadRequest, models.CommonAPIResponse{
			IsError:         true,
			ResponseMessage: "invalid payload format",
			ResponseData:    err,
		})
		return
	}

	orderID := uuid.NewV4().String()

	err := dao.OrdersDao.CreateOrder(orderID, req)
	if err != nil {
		log.Errorf("unable to create order %v", err)
		c.JSON(http.StatusBadRequest, models.CommonAPIResponse{
			IsError:         true,
			ResponseMessage: "unable to create order",
			ResponseData:    err,
		})
		return
	}

	c.JSON(http.StatusOK, models.CommonAPIResponse{
		IsError:         false,
		ResponseMessage: "order created",
		ResponseData:    orderID,
	})
}

func UpdateOrder(c *gin.Context) {
	log := utils.GetLogCtx(c)
	orderID := c.Param("id")
	var req models.UpdateOrderRequest

	if err := c.BindJSON(&req); err != nil {
		log.Errorf("Invalid payload format %v", err)
		c.JSON(http.StatusBadRequest, models.CommonAPIResponse{
			IsError:         true,
			ResponseMessage: "invalid payload format",
			ResponseData:    err,
		})
		return
	}

	var order models.Order
	err := dao.OrdersDao.UpdateOrder(orderID, order, req)
	if err != nil {
		log.Errorf("unable to update order %v", err)
		c.JSON(http.StatusBadRequest, models.CommonAPIResponse{
			IsError:         true,
			ResponseMessage: "unable to update order",
			ResponseData:    err,
		})
		return
	}

	c.JSON(http.StatusOK, models.CommonAPIResponse{
		IsError:         false,
		ResponseMessage: "order status updated successfully",
		ResponseData:    orderID,
	})
}

func GetOrder(c *gin.Context) {
	log := utils.GetLogCtx(c)
	orderID := c.Param("id")

	var order models.Order

	order, err := dao.OrdersDao.GetOrder(orderID, order)
	if err != nil {
		log.Errorf("unable to get order %v", err)
		c.JSON(http.StatusBadRequest, models.CommonAPIResponse{
			IsError:         true,
			ResponseMessage: "unable to get order",
			ResponseData:    err,
		})
		return
	}

	c.JSON(http.StatusOK, models.CommonAPIResponse{
		IsError:         false,
		ResponseMessage: "order fetched successfully",
		ResponseData:    order,
	})
}

func GetOrders(c *gin.Context) {
	log := utils.GetLogCtx(c)
	status := c.Query("status")
	sortBy := c.Query("sortBy")
	sortOrder := c.DefaultQuery("sortOrder", "asc")

	allOrders, err := dao.OrdersDao.GetAllOrders(status, sortBy, sortOrder)
	if err != nil {
		return
	}

	for index, order := range allOrders {
		items, err := dao.OrdersDao.GetOrdersByOrderID(order.ID)
		if err != nil {
			log.Errorf("unable to get order %v", err)
			c.JSON(http.StatusBadRequest, models.CommonAPIResponse{
				IsError:         true,
				ResponseMessage: "unable to get order",
				ResponseData:    err,
			})
			return
		}
		allOrders[index].Items = items
	}

	c.JSON(http.StatusOK, models.CommonAPIResponse{
		IsError:         false,
		ResponseMessage: "order fetched successfully",
		ResponseData:    allOrders,
	})
}
