// Code generated by goctl. DO NOT EDIT.
package types

type IndexReq struct {
}

type IndexResp struct {
	TotalFee          float64        `json:"totalFee" form:"totalFee"`                   // 今日销售金额
	ThisMonthTotalFee float64        `json:"thisMonthTotalFee" form:"thisMonthTotalFee"` // 本月销售金额
	UsedCount         uint           `json:"usedCount" form:"usedCount"`                 // 已使用房间数量
	UnusedCount       uint           `json:"unusedCount" form:"unusedCount"`             // 未使用房间数量
	RoomList          []RoomListItem `json:"roomList" form:"roomList"`                   // 未使用房间数量
}

type RoomListItem struct {
	No         uint    `json:"no"`         // 房号
	Status     uint    `json:"staus"`      // 入住状态
	Balance    float64 `json:"balance"`    // 剩余金额
	SurplusDay int32   `json:"surplusDay"` // 剩余天数

}

type CheckInReq struct {
	No       uint    `json:"no"`       // 房号
	TotalFee float64 `json:"totalFee"` // 金额
	Day      uint    `json:"day"`      // 天数
}

type CheckInResp struct {
}

type CompleteReq struct {
	No uint `json:"no"` // 房号
}

type CompleteResp struct {
}