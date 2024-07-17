package form

type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Email    string `form:"email"`
}
