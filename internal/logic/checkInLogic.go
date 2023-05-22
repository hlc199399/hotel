package logic

import (
	"context"
	"errors"
	"hotel/dict"
	"hotel/internal/model"
	"hotel/internal/svc"
	"hotel/internal/types"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckInLogic {
	return &CheckInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckInLogic) CheckIn(req *types.CheckInReq) (resp *types.CheckInResp, err error) {
	room := l.svcCtx.Db.Room
	r, err := room.WithContext(l.ctx).Where(room.No.Eq(req.No)).Take()
	if err != nil {
		return nil, err
	}
	if r.Status != dict.RoomStatusUnused {
		return nil, err
	}
	roomCondition := l.svcCtx.Db.RoomCondition
	_, err = roomCondition.WithContext(l.ctx).Where(roomCondition.Status.Eq(1)).Take()
	if err == nil {
		return nil, errors.New("该房间被占用")
	}
	// 创建订单
	order := l.svcCtx.Db.Order
	o := model.Order{
		RoomID:   r.ID,
		TotalFee: req.TotalFee,
	}
	err = order.WithContext(l.ctx).Create(&o)
	if err != nil {
		return nil, err
	}
	log.Println("创建订单成功")
	//创建房间信息
	nowTime := time.Now()
	endTime := time.Now().AddDate(0, 0, int(req.Day))
	rc := model.RoomCondition{
		RoomID:      r.ID,
		OrderID:     o.ID,
		No:          r.No,
		CheckInTime: &nowTime,
		EndTime:     &endTime,
		RoomPrice:   r.Price,
		PayFee:      req.TotalFee,
		Balance:     req.TotalFee - r.Price,
		SurplusDay:  int32(req.Day - 1),
		Status:      1,
	}
	err = roomCondition.WithContext(l.ctx).Create(&rc)
	if err != nil {
		return nil, err
	}
	log.Println("创建订单信息成功")
	// 修改房间的状态
	r.Status = dict.RoomStatusUsed
	err = room.WithContext(l.ctx).Save(r)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
