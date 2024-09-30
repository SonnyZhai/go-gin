package models

import "time"

// 自增ID主键
type ID struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
}

// 时间
type Timestamp struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"create_time" form:"create_time"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"update_time" form:"update_time"`
}

// 软删除
type SoftDelete struct {
	DeletedAt time.Time `gorm:"index" json:"delete_time" form:"delete_time"`
}

// 分页
type Page struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}
