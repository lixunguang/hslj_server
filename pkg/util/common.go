package util

import (
	"edu-imp/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetRunPort() string {
	return fmt.Sprintf(":%d", config.Vipper.Get("Server.HttpPort"))
}

func GetMode() string {
	if IsDebug() {
		return gin.DebugMode
	}

	if IsTest() {
		return gin.TestMode
	}

	return gin.ReleaseMode
}

func IsDebug() bool {
	return config.Vipper.Get("Server.RunMode") == gin.DebugMode
}

func IsTest() bool {
	return config.Vipper.Get("Server.RunMode") == gin.TestMode
}

func IsMaster() bool {
	return config.Vipper.Get("Server.ServerType") == "master"
}
