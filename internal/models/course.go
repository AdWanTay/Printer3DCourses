package models

type Course struct {
	ID                   uint    `gorm:"primaryKey"`
	CourseTitle          string  `gorm:"not null" json:"course_title"`
	CourseDescription    string  `gorm:"not null" json:"course_description"`
	NumberOfParticipants string  `json:"number_of_participants"`
	Duration             int     `gorm:"not null"`
	LatexPath            string  `gorm:"unique" json:"latex_path"`
	Price                float32 `gorm:"unique"`
	AuthorName           string  `gorm:"not null" json:"author_name"`
	AuthorTgUsername     string  `gorm:"not null" json:"author_tg_username"`
}
