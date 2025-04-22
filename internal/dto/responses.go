package dto

import "Printer3DCourses/internal/models"

type CoursesPageResponse struct {
	Items []courseResponse `json:"items"`
}

type courseResponse struct {
	ID                   uint    `json:"id"`
	CourseTitle          string  `json:"course_title"`
	ShortDescription     string  `gorm:"not null" json:"short_description"`
	FullDescription      string  `gorm:"not null" json:"full_description"`
	NumberOfParticipants string  `json:"number_of_participants"`
	Duration             int     `json:"duration"`
	Price                float32 `json:"price"`
	IsBought             bool    `json:"is_bought" default:"false"`
}

func NewCoursesPageResponse(items *[]models.Course, paidIndexes map[uint]struct{}) *CoursesPageResponse {
	coursesPageResponse := CoursesPageResponse{}
	courseResponses := make([]courseResponse, 0)
	for _, course := range *items {
		_, exists := paidIndexes[course.ID]

		courseResponses = append(courseResponses, courseResponse{
			ID:                   course.ID,
			CourseTitle:          course.CourseTitle,
			ShortDescription:     course.ShortDescription,
			FullDescription:      course.FullDescription,
			NumberOfParticipants: course.NumberOfParticipants,
			Duration:             course.Duration,
			Price:                course.Price,
			IsBought:             exists,
		})
	}
	coursesPageResponse.Items = courseResponses
	return &coursesPageResponse
}

type ProfilePageResponse struct {
	Items []userCourseResponse `json:"items"`

	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	Patronymic  string `json:"patronymic"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type userCourseResponse struct {
	ID                uint    `json:"id"`
	CourseTitle       string  `json:"course_title"`
	CourseDescription string  `json:"course_description"`
	Progress          float32 `json:"progress"`
}

func NewProfilePageResponse(userCourses *[]models.Course, user *models.User) *ProfilePageResponse {
	items := make([]userCourseResponse, 0, len(*userCourses))

	for _, course := range *userCourses {
		items = append(items, userCourseResponse{
			ID:                course.ID,
			CourseTitle:       course.CourseTitle,
			CourseDescription: course.ShortDescription,
			Progress:          80.0, // можно потом динамически передавать
		})
	}

	return &ProfilePageResponse{
		Items:       items,
		LastName:    user.LastName,
		FirstName:   user.FirstName,
		Patronymic:  user.Patronymic,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}
}

type TestResponse struct {
	Count     uint       `json:"count"`
	Title     string     `json:"title"`
	Questions []question `json:"questions"`
}

type question struct {
	ID           uint     `json:"id"`
	QuestionText string   `json:"question_text"`
	Answers      []answer `json:"answers"`
}

type answer struct {
	ID         uint   `json:"id"`
	AnswerText string `json:"answer_text"`
}

func NewTestResponse(test *models.Test, count int) *TestResponse {
	questionItems := make([]question, 0, len(test.Questions))
	for _, questionItem := range test.Questions {
		answerItems := make([]answer, 0, len(questionItem.Answers))

		for _, answerItem := range questionItem.Answers {
			answerItems = append(answerItems, answer{
				ID:         answerItem.ID,
				AnswerText: answerItem.AnswerText,
			})
		}

		questionItems = append(questionItems, question{
			ID:           questionItem.ID,
			QuestionText: questionItem.QuestionText,
			Answers:      answerItems,
		})
	}

	return &TestResponse{
		Count:     uint(count),
		Title:     test.TestTitle,
		Questions: questionItems,
	}
}
