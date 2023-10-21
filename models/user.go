package models

type Users struct {
	ID                 int64  `gorm:"primaryKey;autoIncrement;index" json:"user_id"`
	Name               string `gorm:"type:varchar(300)" json:"name"`
	Email              string `gorm:"type:varchar(100);index" json:"email"`
	Password           string `gorm:"column:password" json:"-"`
	Address            string `gorm:"type:text" json:"address"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
	IsDeleted          bool   `gorm:"type:boolean" json:"is_deleted"`
	DeletedAt          int64  `gorm:"index" json:"deleted_at"`
	CreatedAtFormatted string `json:"created_at_formatted" gorm:"-"`
	UpdatedAtFormatted string `json:"updated_at_formatted" gorm:"-"`
	DeletedAtFormatted string `json:"deleted_at_formatted" gorm:"-"`

	// Account   Account        `gorm:"foreignKey:UserID"` // Menunjukkan kunci asing
}

type UserCreate struct {
	Name     string `gorm:"type:varchar(300)" json:"name" validate:"required"`
	Email    string `gorm:"type:varchar(100)" json:"email" validate:"required,email,min=6,max=100"`
	Address  string `gorm:"type:text" json:"address" validate:"required"`
	Password string `gorm:"type:varchar(255)" json:"password" validate:"required,min=8"`
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
