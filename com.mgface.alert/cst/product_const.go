package cst

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
