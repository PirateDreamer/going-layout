package entity

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID            int64          `gorm:"primaryKey;column:id;type:bigint unsigned;not null;autoIncrement" json:"id"`                                       // 主键ID
	Username      string         `gorm:"uniqueIndex:uk_username;column:username;type:varchar(50);not null" json:"username"`                                // 用户名
	PasswordHash  string         `gorm:"column:password_hash;type:varchar(255);not null" json:"password_hash"`                                             // 密码哈希
	Salt          *string        `gorm:"column:salt;type:varchar(64)" json:"salt"`                                                                         // 密码盐值
	Email         *string        `gorm:"uniqueIndex:uk_email;column:email;type:varchar(100)" json:"email"`                                                 // 邮箱
	Phone         *string        `gorm:"uniqueIndex:uk_phone;column:phone;type:varchar(20)" json:"phone"`                                                  // 手机号
	Role          int8           `gorm:"column:role;type:tinyint;not null;default:0" json:"role"`                                                          // 角色：0-普通用户 1-管理员 2-其他
	Status        int8           `gorm:"column:status;type:tinyint;not null;default:1" json:"status"`                                                      // 状态：0-禁用 1-正常 2-锁定
	LastLoginTime *time.Time     `gorm:"column:last_login_time;type:datetime" json:"last_login_time"`                                                      // 最后登录时间
	LastLoginIP   *string        `gorm:"column:last_login_ip;type:varchar(45)" json:"last_login_ip"`                                                       // 最后登录IP(支持IPv6)
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"`                             // 创建时间
	UpdatedAt     gorm.DeletedAt `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
}

// TableName 设置表名
func (User) TableName() string {
	return "user"
}
