package service

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/user"
	"automatic-doodle/modules/user/factory"
	"automatic-doodle/modules/user/repository"
	"context"
	"sync"

	"github.com/google/uuid"
)

var (
	moduleOnce sync.Once
	module     UserService
)

func New(ur repository.UserRepository, uf factory.UserFactory) *UserService {

	moduleOnce.Do(func() {
		module = UserService{
			UserRepo:    ur,
			UserFactory: uf,
		}
	})

	return &module

}

type UserService struct {
	UserRepo    repository.UserRepository
	UserFactory factory.UserFactory
}

func (srv *UserService) CreateUser(phoneNumber string, email string, firstName string, lastName string, password string, age int, role user.Role, state user.State, ctx context.Context) (*ent.User, error) {
	userCreate := srv.UserFactory.Create(phoneNumber, email, firstName, lastName, password, role, state)

	return srv.UserRepo.Create(userCreate, ctx)

}

func (srv *UserService) DeleteUser(id uuid.UUID, ctx context.Context) error {
	return srv.UserRepo.DeleteUser(id, ctx)
}
