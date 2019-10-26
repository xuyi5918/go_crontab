package core

import (
	"time"
)

type Log struct {
	LogId 		int

	TaskId		int

	StartTime 	int64
	EndTime 	int64

	LogList		map[string]string
}


// 声明日志记录变量
var logModules map[string](map[int]*Log)

var logList map[string]string

var logTaskRunTime map[string]int64

// 初始化
func init()  {

	logModules  = make(map[string]map[int]*Log)
	logList		= make(map[string]string)
	logTaskRunTime = make(map[string]int64)
	// 启动日志收集方法
	go WriteConsoleLogToDisk()
}

func WriteConsoleLogToDisk()  {

}

func SetConsoleStartTime(taskName string, taskId int) {

	logModules[taskName] = make(map[int]*Log)
	logModules[taskName][taskId] = new(Log)
	logModules[taskName][taskId].TaskId = taskId
	logModules[taskName][taskId].StartTime = time.Now().Unix()

}

func SetConsoleEndTime(taskName string, taskId int) {
	logModules[taskName][taskId].EndTime = time.Now().Unix()
	logTaskRunTime[taskName] = logModules[taskName][taskId].EndTime
}


func GetConsoleLastTime(taskName string) int64{
	return logTaskRunTime[taskName]
}