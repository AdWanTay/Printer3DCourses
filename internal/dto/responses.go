package dto

import "Printer3DCourses/internal/models"

type CoursesPageResponse struct {
	Items []courseResponse `json:"items"`
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

type ProfilePageResponse struct {
	Items []userCourseResponse `json:"items"`

	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
}

type userCourseResponse struct {
	ID                uint    `json:"id"`
	CourseTitle       string  `json:"course_title"`
	CourseDescription string  `json:"course_description"`
	Progress          float32 `json:"progress"`
}

func NewProfilePageResponse(userCourses *[]models.Course, user *models.User) *ProfilePageResponse {
	var userCoursesResponse []userCourseResponse
	for _, course := range *userCourses {
		courseResp := userCourseResponse{
			ID:                course.ID,
			CourseTitle:       course.CourseTitle,
			CourseDescription: course.CourseDescription,
			Progress:          80.0,
		}
		userCoursesResponse = append(userCoursesResponse, courseResp)
	}
	return &ProfilePageResponse{Items: userCoursesResponse,
		LastName: user.LastName, FirstName: user.FirstName,
		Patronymic: user.Patronymic, Email: user.Email}
}
