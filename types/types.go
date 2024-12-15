package types

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenUser struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type Response struct {
	Message string `json:"message"`
}

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
