package user

type IUserService interface {
	Insert(user *RegisterSchema) error
	GetUserByName(userName string) (*User, error)
}
