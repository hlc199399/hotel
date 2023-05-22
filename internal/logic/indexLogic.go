package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/now"
	"hotel/internal/svc"
	"hotel/internal/types"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index() (*types.IndexResp, error) {
	begin := now.New(time.Now()).BeginningOfDay()
	end := now.New(time.Now()).EndOfDay()

	beginMonth := now.New(time.Now()).BeginningOfMonth()
	endMonth := now.New(time.Now()).EndOfMonth()
	log.Printf("++请求时间开始=%s+结束时间=%s+", begin, end)
	roomArr, err := l.svcCtx.Db.Room.Find()
	if err != nil {
		return nil, err
	}

	usedCount := 0
	unusedCount := 0
	var roomListArr []types.RoomListItem
	for _, v := range roomArr {
		item := types.RoomListItem{
			No: v.No,
		}
		if v.Status == 1 {
			item.Status = 1
			unusedCount++
		} else {
			usedCount++
			roomCondition := l.svcCtx.Db.RoomCondition
			rc, err := roomCondition.WithContext(l.ctx).Where(roomCondition.Status.Eq(1)).Last()
			if err != nil {
				break
			}
			item.Balance = rc.Balance
			item.SurplusDay = rc.SurplusDay
			item.Status = 2
			// 如果结束时间是今天，就将返回的状态改成到期
			rcEndTimeStr := rc.EndTime.Format("2006-01-02")
			today := time.Now().Format("2006-01-02")
			if rcEndTimeStr == today {
				item.Status = 3
			}
			// 欠费
			if rc.SurplusDay < 0 {
				item.Status = 4
			}
		}
		roomListArr = append(roomListArr, item)
	}
	oDB := l.svcCtx.Db.Order
	orderArr, err := oDB.Where(oDB.CreatedAt.Between(begin, end)).Find()
	if err != nil {
		return nil, err
	}
	totalFee := 0.0
	for _, v := range orderArr {
		totalFee = totalFee + v.TotalFee
	}
	orderMonthArr, err := oDB.Where(oDB.CreatedAt.Between(beginMonth, endMonth)).Find()
	if err != nil {
		return nil, err
	}
	thisMonthTotalFee := 0.0
	for _, v := range orderMonthArr {
		thisMonthTotalFee = thisMonthTotalFee + v.TotalFee
	}
	resp := &types.IndexResp{
		TotalFee:          totalFee,
		ThisMonthTotalFee: thisMonthTotalFee,
		UsedCount:         uint(usedCount),
		UnusedCount:       uint(unusedCount),
		RoomList:          roomListArr,
	}
	return resp, nil
}

func (l *IndexLogic) AutoUpdateRoomCondition() error {
	fmt.Println("++进入执行+")
	roomCondition := l.svcCtx.Db.RoomCondition
	rcArr, err := roomCondition.WithContext(l.ctx).Where(roomCondition.Status.Eq(1)).Find()
	if err != nil {
		fmt.Println("+++++获取房间信息出错++")
		return nil
	}
	for _, v := range rcArr {

		v.Balance = v.Balance - v.RoomPrice
		v.SurplusDay = v.SurplusDay - 1
		_, err = roomCondition.WithContext(l.ctx).Updates(v)
		if err != nil {
			fmt.Printf("+++++更新房间信息出错+roomNo=%d+err=%v", v.No, err)
			return nil
		}
	}
	return nil
}
