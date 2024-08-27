package user

type User struct {
	ID       string `json:"id" gorm:"type:uuid;primary_key"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
}
