package user

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var IdentityKey = "id"
