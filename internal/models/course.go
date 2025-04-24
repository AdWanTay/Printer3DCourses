package models

type Course struct {
	ID                   uint   `gorm:"primaryKey"`
	CourseTitle          string `gorm:"not null" json:"course_title"`
	ShortDescription     string `gorm:"not null" json:"short_description"`
	FullDescription      string `gorm:"not null" json:"full_description"`
	NumberOfParticipants string `json:"number_of_participants"`
	Duration             int    `gorm:"not null"`
	Price                float32
	AuthorName           string `gorm:"not null" json:"author_name"`
	AuthorTgUsername     string `gorm:"not null" json:"author_tg_username"`
}
