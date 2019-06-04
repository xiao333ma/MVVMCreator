package mvvm

import (
	"MVVMCreator/mvvm/utils"
	"fmt"
	"os/exec"
)

// 一个接口，创建 .h 和 .m
type Creator interface {
	CreateDotHFile()
	CreateDotMFile()
}

func Create(i Creator)()  {
	i.CreateDotHFile()
	i.CreateDotMFile()
}
// 创建 MVVM 文件
func CreateMVVM(prefix string, registerComponentName string) {

	args := utils.GetInstanceArgs()
	args.Prefix = prefix
	args.RegisterComponentName = registerComponentName

	fmt.Printf("努力创建文件中... \n")
	var creator Creator
	creator = new(ViewModelCreator)
	Create(creator)
	creator = new(PipelineCreator)
	Create(creator)
	creator = new(ViewCreator)
	Create(creator)
	creator = new(ViewControllerCreator)
	Create(creator)
	fmt.Printf("🍻 创建完成 \n")
	commandString := "open " + utils.GetPathPrefix(prefix)
	_ = exec.Command("/bin/bash", "-c", commandString).Run()
}