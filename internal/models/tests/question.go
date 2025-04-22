package tests

type Question struct {
	ID           uint   `gorm:"primary_key"`
	QuestionText string `gorm:"type:text"`

	TestID uint `gorm:"not null;OnDelete:CASCADE;"`
	Test   Test `gorm:"foreignKey:TestID;references:ID"`
}
