// Usecase（interface/databaseへのInput Port）
package usecase

import "clean-architecture-tutorial/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
