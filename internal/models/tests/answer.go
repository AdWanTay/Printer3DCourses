package tests

type Answer struct {
	ID         uint   `gorm:"primarykey"`
	AnswerText string `gorm:"type:text"`
	IsCorrect  bool   `gorm:"type:boolean"`

	QuestionID uint     `gorm:"not null;OnDelete:CASCADE;"`
	Question   Question `gorm:"foreignKey:QuestionID;references:ID"`
}
