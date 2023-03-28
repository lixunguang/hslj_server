package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

//判断文件夹/文件是否存在
func IsExist(filePath string) bool {
	_, err := os.Stat(filePath) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(filePath string) bool {
	s, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(filePath string) bool {
	return !IsDir(filePath)
}

//如果文件夹不存在则创建
func CreateDirIfNotExist(filePath string) (bool, error) {
	if IsExist(filePath) && IsDir(filePath) {
		return true, nil
	}

	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return false, err
	}
	return true, nil
}

func IsPythonFile(filePath string) bool {
	return path.Ext(filePath) == ".py"
}

//清理文件夹下的所有文件
func RemoveAllFilesInDir(dirPath string) (bool, error) {
	if !IsExist(dirPath) || !IsDir(dirPath) {
		return false, errors.New("路径错误")
	}

	dirFiles, err := ioutil.ReadDir(dirPath)
	fmt.Println(err, dirFiles)
	for _, d := range dirFiles {
		if err := os.RemoveAll(path.Join([]string{dirPath, d.Name()}...)); err != nil {
			return false, err
		}
	}

	return true, nil
}
