package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"log"
	"order-management-service/api/controllers"
	"order-management-service/api/data"
)

func LoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["request_id"] = params.Keys["X-Request-ID"]
			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["start_time"] = params.TimeStamp.Format("2006/01/02 - 15:04:05")
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var xRequestID string
		if xRequestID = c.Request.Header.Get("X-Request-ID"); xRequestID == "" {
			xRequestID = uuid.NewV4().String()
		}
		c.Set("X-Request-ID", xRequestID)
		log := logrus.WithField("request_id", xRequestID)
		c.Set("log", log)
		c.Next()
	}
}

func Run() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(RequestID())
	router.Use(LoggerMiddleware())

	if err := data.Initialize(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	controllers.Initialize(router)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
