package main

import (
	"console/applicathion/modules/book/conf"
	"console/system/core"
	"fmt"
)

type BookApp struct {
	Name string
}

func (app BookApp) Exec(module string, service string, task int)  {

	fmt.Println("task Id :", task)
	// 执行路由选择
	conf.RouterRun(module, service)

	defer core.Wg.Done();
}

func main()  {
	var c core.App = BookApp{Name:"book"}
	core.Thread(c)
}

