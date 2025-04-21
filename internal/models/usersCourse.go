package models

type UsersCourse struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null;OnDelete:CASCADE;"`
	User   User `gorm:"foreignKey:UserID;references:ID"`

	CourseID uint   `gorm:"not null;OnDelete:CASCADE;"`
	Course   Course `gorm:"foreignKey:CourseID;references:ID"`
}
