package user

type Repository interface {
	Create(*User) error
	FindById(string) (*User, error)
	FindByEmail(string) (*User, error)
}