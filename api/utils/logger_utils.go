package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetLogCtx(ctx *gin.Context) *logrus.Entry {
	log := ctx.Value("log")

	if log == nil {
		logrus.Fatal("Logger is missing in the context") // panics
	}

	return log.(*logrus.Entry)
}
