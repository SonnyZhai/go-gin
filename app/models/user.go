package models

import "strconv"

type User struct {
	ID
	Username string `gorm:"type:varchar(20);not null;uniqueIndex;comment:用户名称" json:"username" form:"username" `
	Password string `gorm:"type:varchar(100);not null;comment:用户名称" json:"-" form:"password"` // 忽略密码字段
	Mobile   string `gorm:"type:varchar(11);not null;uniqueIndex;unique;comment:用户名称" json:"mobile" form:"mobile"`
	Email    string `gorm:"type:varchar(100);not null;uniqueIndex;comment:用户名称" json:"email" form:"email"`
	Timestamp
	SoftDelete
}

// GetUid 获取用户 ID
func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
