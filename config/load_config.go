package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var Vipper *viper.Viper

func init() {
	loadConfig()
}

func loadConfig() {
	vp := viper.New()

	configFilePath := filepath.Join(GetExecDirectory(), "config")
	vp.AddConfigPath(configFilePath)
	//vp.AddConfigPath(getConfigDirPath()) //相对main.go 或 test文件的路径
	vp.SetConfigName("admin-config")

	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Vipper = vp
}

func GetExecDirectory() string {
	file, e := os.Executable()
	if e != nil {
		fmt.Println("error: get Executable file path error")
	}
	path := filepath.Dir(file)

	return path
}

func getConfigDirPath() string {
	configDir := "config"
	// 读取配置文件, 解决跑测试的时候找不到配置文件的问题，最多往上找5层目录
	for i := 0; i < 5; i++ {
		if _, err := os.Stat(configDir); err == nil {
			return configDir
		} else {
			configDir = "../" + configDir
		}
	}

	return configDir
}
