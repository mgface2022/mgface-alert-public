package model

import "time"

// 用户角色常量
const (
	UserRoleUser  = "user"  // 用户
	UserRoleVIP   = "vip"   // VIP
	UserRoleAdmin = "admin" // 管理员
)

// User 结构体表示用户信息
type User struct {
	ID            uint          `gorm:"primaryKey" json:"id"`               // 用户ID，自增主键
	PhoneNumber   string        `gorm:"size:20;unique" json:"phone_number"` // 手机号，非空
	CountryCode   string        `gorm:"size:10" json:"country_code"`        // 国家码，非空
	Username      string        `gorm:"size:50;not null;unique" json:"username"`
	Password      string        `gorm:"size:100;not null" json:"password"`
	Email         string        `gorm:"size:100;unique" json:"email"`
	Avatar        string        `gorm:"size:255" json:"avatar"`                      // 头像URL
	Role          string        `gorm:"size:20;not null;default:'user'" json:"role"` // 用户角色
	Status        string        `gorm:"size:20;not null" json:"status"`              // 状态，非空
	CreatedAt     time.Time     `gorm:"autoCreateTime" json:"created_at"`            // 创建时间，自动创建
	UpdatedAt     time.Time     `gorm:"autoUpdateTime" json:"updated_at"`            // 更新时间，自动更新
	LastLogin     time.Time     `json:"last_login"`                                  // 最后登录时间
	UserVIPStatus UserVIPStatus `gorm:"foreignKey:UserID" json:"user_vip_status"`    // 添加外键关联
}

// UserInfo 返回给前端的用户信息结构体
type UserInfo struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
}

// ToUserInfo 将 User 转换为 UserInfo
func (u *User) ToUserInfo() *UserInfo {
	return &UserInfo{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Avatar:    u.Avatar,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		LastLogin: u.LastLogin,
	}
}
