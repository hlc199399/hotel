package timer

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"hotel/internal/logic"
	"hotel/internal/svc"
)

const (
	//refreshTime = "0 0 1 * * ?"
	refreshTime = "@every 1m"
)

var cronTask *cron.Cron

var serverCtx *svc.ServiceContext

func InitTask(s *svc.ServiceContext) {
	serverCtx = s
	//清除定时器
	if cronTask != nil {
		cronTask.Stop()
		var entries = cronTask.Entries()
		for _, entry := range entries {
			cronTask.Remove(entry.ID)
		}
		cronTask = nil
	}

	//创建定时器
	if cronTask == nil {
		cronTask = cron.New()
	}

	//添加定时任务

	_, err := cronTask.AddFunc(refreshTime, handler)
	if err != nil {
		panic(err)
	}

	//检查是否需要启动任务
	if len(cronTask.Entries()) > 0 {
		cronTask.Start()
	} else {
		cronTask = nil
	}
}

func handler() {
	fmt.Println("++执行")
	l := logic.NewIndexLogic(context.Background(), serverCtx)
	l.AutoUpdateRoomCondition()
}
