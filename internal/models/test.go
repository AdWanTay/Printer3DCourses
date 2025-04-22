package models

type Test struct {
	ID        uint   `gorm:"primaryKey"`
	TestTitle string `json:"test_title"`

	CourseID uint   `gorm:"not null;OnDelete:CASCADE;"`
	Course   Course `gorm:"foreignKey:CourseID;references:ID"`

	Question []Question `gorm:"foreignKey:TestID;references:ID"`
}

type Question struct {
	ID           uint   `gorm:"primaryKey"`
	QuestionText string `gorm:"type:text"`

	TestID uint `gorm:"not null;OnDelete:CASCADE;"`
	Test   Test `gorm:"foreignKey:TestID;references:ID"`
}

type Answer struct {
	ID         uint   `gorm:"primaryKey"`
	AnswerText string `gorm:"type:text"`
	IsCorrect  bool   `gorm:"type:boolean"`

	QuestionID uint     `gorm:"not null;OnDelete:CASCADE;"`
	Question   Question `gorm:"foreignKey:QuestionID;references:ID"`
}

type TestScore struct {
	ID uint64 `gorm:"primaryKey"`

	TestID uint `gorm:"not null;OnDelete:CASCADE;"`
	Test   Test `gorm:"foreignKey:TestID;references:ID"`

	UserID uint `gorm:"not null;OnDelete:CASCADE;"`
	User   User `gorm:"foreignKey:UserID;references:ID"`
}
