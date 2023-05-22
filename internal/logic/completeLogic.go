package logic

import (
	"context"
	"errors"
	"hotel/dict"
	"hotel/internal/svc"
	"hotel/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteLogic {
	return &CompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompleteLogic) Complete(req *types.CompleteReq) (resp *types.CompleteResp, err error) {
	room := l.svcCtx.Db.Room
	r, err := room.WithContext(l.ctx).Where(room.No.Eq(req.No)).Take()
	if err != nil {
		return nil, err
	}
	if r.Status == dict.RoomStatusUnused {
		return nil, errors.New("该房间未使用")
	}
	// 将房间信息表的状态改为已完成
	roomCondition := l.svcCtx.Db.RoomCondition
	rc, err := roomCondition.WithContext(l.ctx).Where(roomCondition.Status.Eq(1)).Last()
	if err != nil {
		return nil, errors.New("未找到使用中的房间")
	}
	rc.Status = 2
	err = roomCondition.WithContext(l.ctx).Save(rc)
	if err != nil {
		return nil, err
	}
	// 将房间状态改为未使用
	r.Status = 1
	err = room.WithContext(l.ctx).Save(r)
	if err != nil {
		return nil, err
	}
	return
}
