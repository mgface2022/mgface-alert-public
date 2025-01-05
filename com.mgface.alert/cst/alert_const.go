package cst

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
