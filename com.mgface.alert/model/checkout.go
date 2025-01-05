package model

import (
	"errors"
	"strconv"
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
