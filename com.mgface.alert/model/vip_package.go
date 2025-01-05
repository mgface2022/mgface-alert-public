package model

import (
	"errors"
	"strconv"
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

// CheckoutRequest VIP结账请求结构体
type CheckoutRequest struct {
	ProductID     string `json:"product_id" binding:"required"`
	Amount        string `json:"amount" binding:"required"`
	CouponCode    string `json:"coupon_code"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

// ToUint 将ProductID转换为uint
func (r *CheckoutRequest) ToUint() (uint, error) {
	id, err := strconv.ParseUint(r.ProductID, 10, 64)
	if err != nil {
		return 0, errors.New("vip_package.product_id.invalid")
	}
	return uint(id), nil
}

// ToFloat64 将Amount转换为float64
func (r *CheckoutRequest) ToFloat64() (float64, error) {
	amount, err := strconv.ParseFloat(r.Amount, 64)
	if err != nil {
		return 0, errors.New("vip_package.amount.invalid")
	}
	return amount, nil
}

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
