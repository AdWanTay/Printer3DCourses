package dto

type RegistrationResponse struct {
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	Patronymic  string `json:"patronymic"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}
