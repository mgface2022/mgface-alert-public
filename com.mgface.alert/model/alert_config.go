package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// 价格条件常量
const (
	PriceConditionAbove = "above" // 高于
	PriceConditionBelow = "below" // 低于
)

// 通知方式常量
const (
	NotifyTypeSMS      = "sms"      // 短信
	NotifyTypePhone    = "phone"    // 电话
	NotifyTypeEmail    = "email"    // 邮件
	NotifyTypeCheckOut = "recharge" // 充值
)

// 告警频率常量
const (
	AlertFrequency30MIN = "30m" // 30分钟
	AlertFrequency1H    = "1h"  // 1小时
	AlertFrequency2H    = "2h"  // 2小时
	AlertFrequency4H    = "4h"  // 4小时
	AlertFrequency8H    = "8h"  // 8小时
	AlertFrequency12H   = "12h" // 12小时
	AlertFrequency24H   = "24h" // 24小时
)

// 告警间隔时间常量
const (
	AlertInterval5Min  = "5m"  // 5分钟
	AlertInterval15Min = "15m" // 15分钟
	AlertInterval30Min = "30m" // 30分钟
)

// NotifyTypes 通知方式数组类型
type NotifyTypes []string

// Value 实现 driver.Valuer 接口
func (n NotifyTypes) Value() (driver.Value, error) {
	return json.Marshal(n)
}

// Scan 实现 sql.Scanner 接口
func (n *NotifyTypes) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), n)
}

// AlertConfig 告警配置结构体
type AlertConfig struct {
	ID             uint        `gorm:"primaryKey" json:"id"`                                 // 配置ID
	CoinCode       string      `gorm:"size:20;not null" json:"coin_code"`                    // 币种代码
	PriceCondition string      `gorm:"size:20;not null" json:"price_condition"`              // 价格条件
	TargetPrice    float64     `gorm:"type:decimal(20,8);not null" json:"target_price"`      // 目标价格
	NotifyTypes    NotifyTypes `gorm:"type:json;not null" json:"notify_types"`               // 通知方式（支持多选）
	AlertFrequency string      `gorm:"size:20;not null;default:'1h'" json:"alert_frequency"` // 告警频率
	AlertCount     int         `gorm:"not null;default:1" json:"alert_count"`                // 告警次数
	AlertInterval  string      `gorm:"size:20;not null;default:'5m'" json:"alert_interval"`  // 告警间隔时间
	User           User        `gorm:"foreignKey:UserID" json:"-"`                           // 关联的用户信息
	UserID         uint        `gorm:"not null" json:"user_id"`                              // 用户ID
	Status         string      `gorm:"size:20;not null;default:'enabled'" json:"status"`     // 状态
	CreatedAt      time.Time   `gorm:"autoCreateTime" json:"created_at"`                     // 创建时间
	UpdatedAt      time.Time   `gorm:"autoUpdateTime" json:"updated_at"`                     // 更新时间
}

// ValidateNotifyTypes 验证通知方式是否合法
func (a *AlertConfig) ValidateNotifyTypes() bool {
	validTypes := map[string]bool{
		NotifyTypeSMS:   true,
		NotifyTypePhone: true,
		NotifyTypeEmail: true,
	}

	for _, t := range a.NotifyTypes {
		if !validTypes[t] {
			return false
		}
	}
	return true
}

// ValidateAlertFrequency 验证告警频率是否合法
func (a *AlertConfig) ValidateAlertFrequency() bool {
	validFrequencies := map[string]bool{
		AlertFrequency30MIN: true,
		AlertFrequency1H:    true,
		AlertFrequency2H:    true,
		AlertFrequency4H:    true,
		AlertFrequency8H:    true,
		AlertFrequency12H:   true,
		AlertFrequency24H:   true,
	}
	return validFrequencies[a.AlertFrequency]
}

// ValidateAlertInterval 验证告警间隔时间是否合法
func (a *AlertConfig) ValidateAlertInterval() bool {
	validIntervals := map[string]bool{
		AlertInterval5Min:  true,
		AlertInterval15Min: true,
		AlertInterval30Min: true,
	}
	return validIntervals[a.AlertInterval]
}

// ValidateAlertCount 验证告警次数是否合法
func (a *AlertConfig) ValidateAlertCount() bool {
	return a.AlertCount >= 1 && a.AlertCount <= 3
}

// Validate 验证告警配置是否合法
func (a *AlertConfig) Validate() error {
	if !a.ValidateNotifyTypes() {
		return errors.New("alert_config.notify_types.invalid")
	}
	if !a.ValidateAlertFrequency() {
		return errors.New("alert_config.frequency.invalid")
	}
	if !a.ValidateAlertInterval() {
		return errors.New("alert_config.interval.invalid")
	}
	if !a.ValidateAlertCount() {
		return errors.New("alert_config.count.invalid")
	}
	return nil
}
