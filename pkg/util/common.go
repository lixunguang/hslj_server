package util

import (
<<<<<<< HEAD
	"hslj/config"
=======
>>>>>>> 688a4df92fb5de6d3a3c38567476cf81c98e2521
	"fmt"
	"github.com/gin-gonic/gin"
	"hslj/config"
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
