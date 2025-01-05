package model

import (
	"time"
)

// UserVIPStatus 用户VIP状态结构体
type UserVIPStatus struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	UserID         uint       `gorm:"not null;uniqueIndex" json:"user_id"`           // 用户ID
	IsVIP          bool       `gorm:"column:is_vip" json:"is_vip"`                   // 是否是VIP
	VIPExpireTime  *time.Time `gorm:"column:vip_expire_time" json:"vip_expire_time"` // 修改这里，使用column标签
	CoinCount      int        `gorm:"not null" json:"coin_count"`                    // 可选币种数量
	PhoneCallCount int        `gorm:"not null" json:"phone_call_count"`              // 可用电话通知次数
	EmailCount     int        `gorm:"not null" json:"email_count"`                   // 可用邮箱通知次数
	SMSCount       int        `gorm:"not null" json:"sms_count"`                     // 可用短信通知次数
	CreatedAt      time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

func (UserVIPStatus) TableName() string {
	return "user_vip_status"
}
