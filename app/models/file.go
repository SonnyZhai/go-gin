package models

type File struct {
	ID
	UserId   int    `json:"user_id" gorm:"index;not null;comment:用户ID"`
	FileType string `json:"media_type" gorm:"size:20;index;not null;comment:文件"`
	FileName string `json:"file_name" gorm:"size:255;not null;comment:文件名"`
	FileSize int64  `json:"file_size" gorm:"not null;comment:文件大小"`
	FilePath string `json:"file_path" gorm:"size:255;comment:文件路径"`
	FileUrl  string `json:"file_url" gorm:"size:255;not null;comment:文件URL"`
	Timestamp
}
