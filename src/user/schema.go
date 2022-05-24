package user

type RegisterSchema struct {
	UserName  string `json:"username" binding:"required"`
	FirstName string `json:"firstname" binding:"alphanum"`
	LastName  string `json:"lastname" binding:"alphanum"`
	Password  string `json:"password" binding:"required"`
}
