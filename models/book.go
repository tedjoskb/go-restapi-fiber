package models

type Book struct {
	Id                 int64  `gorm:"primaryKey" json:"id"`
	Title              string `gorm:"type:varchar(300)" json:"title"`
	Description        string `gorm:"type:text" json:"description"`
	Author             string `gorm:"type:varchar(300)" json:"author"`
	Cover              string `gorm:"type:date" json:"cover"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
	IsDeleted          bool   `gorm:"type:boolean" json:"is_deleted"`
	DeletedAt          int64  `gorm:"index" json:"deleted_at"`
	CreatedAtFormatted string `json:"created_at_formatted" gorm:"-"`
	UpdatedAtFormatted string `json:"updated_at_formatted" gorm:"-"`
	DeletedAtFormatted string `json:"deleted_at_formatted" gorm:"-"`
}
