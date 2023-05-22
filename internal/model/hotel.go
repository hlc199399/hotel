package model

import (
	"gorm.io/gorm"
	"time"
)

// User 用户
type User struct {
	gorm.Model
	Name   string `json:"name"`   // 名字
	IDCard string `json:"idCard"` // 身份证
	Age    uint   `json:"age"`    // 年龄
	Gender uint   `json:"gender"` // 性别 1: 男 2：女 0 未知
}

// Room 房间
type Room struct {
	gorm.Model
	No       uint    `json:"no"`       // 房号
	RoomType uint    `json:"roomType"` // 房间类型 1.临时，2.长期
	Status   uint    `json:"status"`   // 房间状态 1.未入住、2.已入住、3. 今天到期
	Price    float64 `json:"price"`    // 价格
}

type Order struct {
	gorm.Model
	UserID   uint    `json:"userID"`   // 用户ID
	RoomID   uint    `json:"roomID"`   // 房间ID
	TotalFee float64 `json:"totalFee"` // 金额
}

type RoomCondition struct {
	gorm.Model
	RoomID      uint       `json:"roomID"`      // 房间ID
	OrderID     uint       `json:"orderID"`     // 订单ID
	No          uint       `json:"no"`          // 房间号
	CheckInTime *time.Time `json:"checkInTime"` // 入住时间
	EndTime     *time.Time `json:"endTime"`     // 结束时间
	RoomPrice   float64    `json:"roomPrice"`   // 房间价格
	PayFee      float64    `json:"payFee"`      // 支付金额
	Balance     float64    `json:"balance"`     // 剩余金额
	SurplusDay  int32      `json:"surplusDay"`  // 剩余天数
	Status      uint       `json:"status"`      // 状态   1. 未完成 2. 已完成
}
