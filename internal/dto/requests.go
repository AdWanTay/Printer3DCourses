package dto

type RegistrationRequest struct {
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	Patronymic  string `json:"patronymic"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StarterKitRequest struct {
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type TestAnswersRequest struct {
	TestId uint          `json:"test_id"`
	Result map[uint]uint `json:"result"`
}
