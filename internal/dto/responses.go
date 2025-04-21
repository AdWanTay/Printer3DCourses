package dto

import "Printer3DCourses/internal/models"

type CoursesPageResponse struct {
	Items []courseResponse
}

type courseResponse struct {
	ID                   uint    `json:"id"`
	CourseTitle          string  `json:"course_title"`
	CourseDescription    string  `json:"course_description"`
	NumberOfParticipants string  `json:"number_of_participants"`
	Duration             int     `json:"duration"`
	Price                float32 `json:"price"`
}

func NewCoursesPageResponse(items *[]models.Course) *CoursesPageResponse {
	coursesPageResponse := CoursesPageResponse{}
	courseResponses := make([]courseResponse, 0)
	for _, course := range *items {
		courseResponses = append(courseResponses, courseResponse{
			ID:                   course.ID,
			CourseTitle:          course.CourseTitle,
			CourseDescription:    course.CourseDescription,
			NumberOfParticipants: course.NumberOfParticipants,
			Duration:             course.Duration,
			Price:                course.Price,
		})
	}
	coursesPageResponse.Items = courseResponses
	return &coursesPageResponse
}
