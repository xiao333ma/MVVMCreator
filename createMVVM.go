package main

import (
	"./mvvm"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) >= 3 {
		prefix := "IGC" + os.Args[1]
		registerComponentName := os.Args[2]
		mvvm.CreateMVVM(prefix, registerComponentName)
	} else  {
		fmt.Print(`
该程序用于创建 MVVM 相关类的文件
需要输入两个参数

1. 请输入要创建的 MVVM 套件的前缀名称
   例如如果想创建 IGCHome 相关的 VM View VC Pipeline，则输入 Home
2. 请输入 VC 注册的组件名称，例如 igetcool://your/path/to/register/this/viewController
		`)
	}


}
