package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var(
	Wg sync.WaitGroup
)


type App interface {
	Exec(module string, service string, task int)
}

type thread struct {
	Thread_num 	int
	Exec_time 	int
	Modules 	string
	Service 	string
}

type initConf struct {
	Thread map[string] thread
}

/**
	读取项目配置
 */
func getInitConf() initConf {
	getFile, getError:= ioutil.ReadFile("./conf/init.json")
	if getError != nil {
		fmt.Println(getError)
	}

	var response initConf
	unmarshalError := json.Unmarshal(getFile, &response)
	if unmarshalError != nil {
		fmt.Println(unmarshalError)
	}

	return  response
}


func Produce(ch chan thread, conf initConf) int {

	var threadNum = 0; // default 0

	for _, v:= range conf.Thread {

		for num:=0; v.Thread_num > num; num++ {
				ch <- v;
				threadNum ++
		}

	}

	return threadNum
}

func Thread(app App) {

		conf := getInitConf()

		// 生产
		ch := make(chan thread, 10000)
		var threadNum int;
		Wg.Add(1)
		go func() {
			threadNum = Produce(ch, conf)
			Wg.Done()
		}()
		Wg.Wait()


		// 消费线程
		for task:=1; task <= threadNum; task++ {
			taskRow := <- ch
			Wg.Add(1)
			SetConsoleStartTime(taskRow.Service, task)
			go app.Exec(taskRow.Modules, taskRow.Service, task)
			SetConsoleEndTime(taskRow.Service, task)
		}

		Wg.Wait()
}