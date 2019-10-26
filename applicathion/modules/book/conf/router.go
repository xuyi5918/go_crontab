package conf

import (
	"console/applicathion/modules/book/spider"
	"console/applicathion/modules/book/web"
	"fmt"
	"net"
)

var router map[string]func()

func init()  {
	router = make(map[string]func())

	// spider group router
	router["spider:chuangshi"] = func() { spider.Chuangshi() }
	router["spider:qidian"] = func() { spider.Qidian() }

	// web group router
	router["web:users"] = func() {
		web.UsersReadLog()
	}

}

func RouterRun(dir string, file string) interface{} {

	var sign string
	sign = fmt.Sprintf("%s:%s", dir, file)
	router[sign]()

	return net.Interface{}
}
