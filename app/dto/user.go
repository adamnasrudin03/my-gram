package dto

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	Username string `json:"username" `
	Email    string `json:"email" `
	Password string `json:"password"`
	Age      uint64 `json:"age"`
}
