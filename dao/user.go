package dao

type CreateUser struct {
	Name     string `json:"name" example:"admin"`
	Password string `json:"password" example:"123456"`
}

type DeleteUser struct {
	ID uint `json:"id"`
}

type UpdateUser struct {
	ID       string `json:"id"`
	Name     string `json:"name" example:"admin"`
	Password string `json:"password" example:"123456"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type FindUserByNameAndPwd struct {
	Name     string `json:"name" example:"admin"`
	Password string `json:"password" example:"123456"`
}
