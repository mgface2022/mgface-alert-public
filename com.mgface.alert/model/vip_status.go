package model

import (
	"gorm.io/gorm"
	"time"
)

// VIP套餐类型常量
const (
	VIPTypeMonthly   = "monthly"   // 月度
	VIPTypeQuarterly = "quarterly" // 季度
	VIPTypeYearly    = "yearly"    // 年度
)

// 支付状态常量
const (
	PaymentStatusSuccess   = "success"   // 成功|订单已支付成功
	PaymentStatusPending   = "pending"   // 待支付|订单已创建，等待支付
	PaymentStatusFailed    = "failed"    // 支付失败|支付过程中出现错误
	PaymentStatusCancelled = "cancelled" // 已取消|订单已取消
	PaymentStatusRefunded  = "refunded"  // 退款|订单已退款
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

// GetUserVIPStatus 获取用户的VIP状态
func GetUserVIPStatus(db *gorm.DB, userID uint) (*UserVIPStatus, error) {
	type Result struct {
		ID             uint       `gorm:"column:id"`
		UserID         uint       `gorm:"column:user_id"`
		IsVIP          int        `gorm:"column:is_vip"`          // 使用int接收0/1
		VIPExpireTime  *time.Time `gorm:"column:vip_expire_time"` // 明确指定列名
		SMSCount       int        `gorm:"column:sms_count"`
		EmailCount     int        `gorm:"column:email_count"`
		PhoneCallCount int        `gorm:"column:phone_call_count"`
	}

	var result Result
	err := db.Table("user_vip_status").
		Select("id, user_id, is_vip, vip_expire_time, sms_count, email_count, phone_call_count").
		Where("user_id = ?", userID).
		First(&result).Error
	if err != nil {
		return nil, err
	}

	// 转换为UserVIPStatus结构体
	status := &UserVIPStatus{
		ID:             result.ID,
		UserID:         result.UserID,
		IsVIP:          result.IsVIP == 1, // 将0/1转换为false/true
		VIPExpireTime:  result.VIPExpireTime,
		SMSCount:       result.SMSCount,
		EmailCount:     result.EmailCount,
		PhoneCallCount: result.PhoneCallCount,
	}

	return status, nil
}
