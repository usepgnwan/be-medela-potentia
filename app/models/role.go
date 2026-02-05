package models

type UserRole struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;"json:"id" swaggerignore:"true"`
	Deskripsi string `gorm:"size:255" json:"deskripsi" validate:"required"`
}

func (UserRole) TableName() string {
	return "tbl_role"
}
