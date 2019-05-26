package utils

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"sync"
	"time"
)

type CreateArgs struct {
	Prefix string
	RegisterComponentName string
}

var m *CreateArgs
var once sync.Once

// 单利，用于保存命令行参数
func GetInstanceArgs() *CreateArgs {
	once.Do(func() {
		m = &CreateArgs {}
	})
	return m
}

// 文件类型
type FileType int

const (
	FileTypeOfViewModel FileType = iota
	FileTypeOfViewController
	FileTypeOfView
	FileTypeOfPipeLine
)

//
var fileTypeMap = map[FileType]string{
	FileTypeOfViewModel:      "ViewModel",
	FileTypeOfViewController: "ViewController",
	FileTypeOfView:           "View",
	FileTypeOfPipeLine:       "Pipeline",
}

// 获取工程名称
func getPorjName() string {
	return "DDProject"
}

// 获取用户名，电脑的登录名
func getUserName() string {
	u, _ := user.Current()

	return u.Username
}

// 获取当前时间
func getCurrentTime() string  {
	return time.Now().Format("02/01/2006")
}

// 获取当前的年份
func getCurrentYear() string {
	return time.Now().Format("2006")
}

// 获取当前的路径，用于结束后打开文件夹
func GetPathPrefix(prefix string) string {
	return "/Users/" + getUserName() + "/Desktop/CreateMVVM/" + prefix
}

// 获取文件的全部路径
func getFullPath(fileType FileType, prefix string) string {
	return GetPathPrefix(prefix) + "/" + prefix + fileTypeMap[fileType]
}

// 获取文件名字
func getFileFullName(fileType FileType, prefix string, header bool) string {
	suffix := ""
	if header {
		suffix = ".h"
	} else  {
		suffix = ".m"
	}
	return prefix + fileTypeMap[fileType] + suffix
}

// 写入文件，替换对应的字符串，并写入
func WriteFile(fileType FileType, header bool, content string)  {
	prefix := GetInstanceArgs().Prefix
	content = replace(content, prefix)
	writeFile(fileType, prefix, header, content)
}

// 替换对应的字符串
func replace(str string, prefix string) string {
	str = strings.Replace(str, "__prefix__", prefix, -1)
	str = strings.Replace(str, "__user__", getUserName(), -1)
	str = strings.Replace(str, "__time__", getCurrentTime(), -1)
	str = strings.Replace(str, "__year__", getCurrentYear(), -1)
	str = strings.Replace(str, "__proj__", getPorjName(), -1)
	str = strings.Replace(str, "__component__name__", GetInstanceArgs().RegisterComponentName, -1)
	return str
}

// 真正写入文件
func writeFile(fileType FileType, prefix string, header bool, content string)  {

	fullPath := getFullPath(fileType, prefix)
	fileFullName := getFileFullName(fileType, prefix, header)
	createFolderAndWriteToFile(fullPath, fileFullName, content)
}

// 创建文件夹，然后写入文件
func createFolderAndWriteToFile(fullPath string, fileFullName string, content string)  {
	ok := createFolderIfNotExists(fullPath)
	if ok {
		f, err := os.OpenFile(fullPath + "/" + fileFullName, os.O_CREATE | os.O_WRONLY, os.ModePerm)
		if err == nil {
			 _, writeErr := f.WriteString(content)
			if writeErr == nil {
				fmt.Printf("🎉 " + fileFullName + " 创建成功\r\n")
			} else  {
				fmt.Printf("💥 " + fileFullName + " 创建失败\r\n")
			}
		}
	}
}

// 判断文件夹是否存在，如果不存在的话，就创建
func createFolderIfNotExists(fullPath string) bool {
	_, err := os.Stat(fullPath)
	if err != nil {
		mkdirError := os.MkdirAll(fullPath, os.ModePerm)
		if mkdirError != nil {
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}