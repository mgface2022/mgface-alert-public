package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	. "github.com/mgface2022/mgface-alert-public/com.mgface.alert/cst"
	"time"
)

func GetBeforeTime(alertFrequency string, currentTime time.Time) time.Time {
	var startTime time.Time
	switch alertFrequency {
	case AlertFrequency30MIN:
		startTime = currentTime.Add(-30 * time.Minute)
	case AlertFrequency1H:
		startTime = currentTime.Add(-time.Duration(1) * time.Hour)
	case AlertFrequency2H:
		startTime = currentTime.Add(-time.Duration(2) * time.Hour)
	case AlertFrequency4H:
		startTime = currentTime.Add(-time.Duration(4) * time.Hour)
	case AlertFrequency8H:
		startTime = currentTime.Add(-time.Duration(8) * time.Hour)
	case AlertFrequency12H:
		startTime = currentTime.Add(-time.Duration(12) * time.Hour)
	case AlertFrequency24H:
		startTime = currentTime.Add(-time.Duration(24) * time.Hour)
	}
	return startTime
}

func GetInterval(alertInterval string, lastNotifyTime time.Time) time.Time {
	var nextAllowedTime time.Time
	switch alertInterval {
	case AlertInterval5Min:
		nextAllowedTime = lastNotifyTime.Add(time.Duration(5) * time.Minute)
	case AlertInterval15Min:
		nextAllowedTime = lastNotifyTime.Add(time.Duration(15) * time.Minute)
	case AlertInterval30Min:
		nextAllowedTime = lastNotifyTime.Add(time.Duration(30) * time.Minute)
	}
	return nextAllowedTime
}

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
