package user

type UserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}
