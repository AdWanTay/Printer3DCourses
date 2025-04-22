package tests

type Question struct {
	ID           uint   `gorm:"primaryKey"`
	QuestionText string `gorm:"type:text"`

	TestID uint `gorm:"not null;OnDelete:CASCADE;"`
	Test   Test `gorm:"foreignKey:TestID;references:ID"`
}
