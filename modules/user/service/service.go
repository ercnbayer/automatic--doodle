package service

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/user"
	"context"
	"sync"

	"github.com/google/uuid"
)

var (
	moduleOnce sync.Once
	module     UserService
)

func New(ur UserRepository, uf UserFactory, ff FileFactory, fr FileRepository) *UserService {

	moduleOnce.Do(func() {
		module = UserService{
			userRepository: ur,
			userFactory:    uf,
			fileFactory:    ff,
			fileRepository: fr,
		}
	})

	return &module

}

type UserService struct {
	userRepository UserRepository
	userFactory    UserFactory
	fileFactory    FileFactory
	fileRepository FileRepository
	logger         Logger
}

func (srv *UserService) CreateUser(phoneNumber string, email string, firstName string, lastName string, password string, role user.Role, state user.State, ctx context.Context) (*ent.User, error) {
	userCreate := srv.userFactory.Create(phoneNumber, email, firstName, lastName, password, role, state)

	return srv.userRepository.Create(userCreate, ctx)

}

func (srv *UserService) DeleteUser(id uuid.UUID, ctx context.Context) error {
	return srv.userRepository.DeleteUser(id, ctx)
}
