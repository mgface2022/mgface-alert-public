package tools

import (
	"os"
	"strings"
)

// Above @txPrice
func Above(txPrice, settingPrice float64) bool {
	return txPrice > settingPrice
}

func Below(txPrice, settingPrice float64) bool {
	return txPrice < settingPrice
}

// IsProductionMode 判断是否是生产环境
func IsProductionMode() bool {
	mode := strings.ToLower(os.Getenv("RUN_MODE"))
	return mode == "prod" || mode == "production"
}
