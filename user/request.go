package user

type UserRequest struct {
	Email    string `json:"email" validate="required,mail"`
	Password string `json:"password" validate="required"`
	Name     string `json:"name" validate="required"`
}
