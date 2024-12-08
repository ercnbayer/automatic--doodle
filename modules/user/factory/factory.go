package factory

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/user"
	"sync"
	"time"

	"github.com/google/uuid"
)

type UserFactory struct {
	db         *ent.Client
	encryption Encryption
}

func (f *UserFactory) Create(
	phoneNumber,
	email,
	firstName,
	lastName,

	password string,
	role user.Role,
	state user.State,
) *ent.UserCreate {
	salt := f.encryption.GenerateSalt(16)
	hash := f.encryption.HashPassword(password, salt)
	now := time.Now()

	item := f.db.User.Create().
		SetID(uuid.New()).
		SetEmail(email).
		SetPhoneNumber(phoneNumber).
		SetLastName(lastName).
		SetPasswordSalt(salt).
		SetPasswordHash(hash).
		SetRole(role).
		SetState(state).
		SetCreatedAt(now).
		SetUpdatedAt(now)

	return item
}

var (
	module     UserFactory
	moduleOnce sync.Once
)

func New(
	db *ent.Client,
	//encryption Encryption,
) *UserFactory {
	moduleOnce.Do(func() {
		module = UserFactory{
			db: db,
		}
	})

	return &module
}
