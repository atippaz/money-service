package interfaces

type AuthRegisterRequest struct {
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password"`
	Profile     string `json:"profile"`
}
