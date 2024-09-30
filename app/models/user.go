package models

type User struct {
	ID
	Username string `gorm:"type:varchar(20);not null;unique" json:"username" form:"username"`
	Password string `gorm:"type:varchar(100);not null" json:"password" form:"password"`
	Mobile   string `gorm:"type:varchar(11);not null;unique" json:"mobile" form:"mobile"`
	Email    string `gorm:"type:varchar(100);not null;unique" json:"email" form:"email"`
	Timestamp
	SoftDelete
}
