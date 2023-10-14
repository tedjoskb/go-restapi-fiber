package models

type Users struct {
	ID                 int64  `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Name               string `gorm:"type:varchar(300)" json:"name"`
	Email              string `gorm:"type:varchar(100)" json:"email"`
	Address            string `gorm:"type:text" json:"address"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
	IsDeleted          bool   `gorm:"type:boolean" json:"is_deleted"`
	DeletedAt          int64  `gorm:"index" json:"deleted_at"`
	CreatedAtFormatted string `json:"created_at_formatted" gorm:"-"`
	UpdatedAtFormatted string `json:"updated_at_formatted" gorm:"-"`
	DeletedAtFormatted string `json:"deleted_at_formatted" gorm:"-"`

	// Account   Account        `gorm:"foreignKey:UserID"` // Menunjukkan kunci asing

	//  atribut inline struct untuk format waktu agar tidak termigrasi ke database
	// TimeFormats struct {
	// 	CreatedAtFormatted string `json:"created_at_formatted"`
	// 	UpdatedAtFormatted string `json:"updated_at_formatted"`
	// 	DeletedAtFormatted string `json:"deleted_at_formatted"`
	// } `gorm:"-"`
}

type UserCreate struct {
	Name    string `gorm:"type:varchar(300)" json:"name" validate:"required"`
	Email   string `gorm:"type:varchar(100)" json:"email" validate:"required"`
	Address string `gorm:"type:text" json:"address" validate:"required"`
}

type UserUpdate struct {
	Name    string `gorm:"type:varchar(300)" json:"name" validate:"required"`
	Address string `gorm:"type:text" json:"address" validate:"required"`
	Email   string `gorm:"type:varchar(100)" json:"email" validate:"required"`
}

type UserUpdateBulky struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Name      string `gorm:"type:varchar(300)" json:"name"`
	Email     string `gorm:"type:varchar(100)" json:"email"`
	Address   string `gorm:"type:text" json:"address"`
	IsDeleted bool   `gorm:"type:boolean" json:"is_deleted"`
}

type SoftDeleteUser struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"user_id"`
}

type UserEmailRequest struct {
	Email string `gorm:"type:varchar(100)" json:"email" validate:"required"`
}
