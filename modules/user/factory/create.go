package factory

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/user"
	"time"

	"github.com/google/uuid"
)

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
		SetFirstName(firstName).
		SetPasswordSalt(salt).
		SetPasswordHash(hash).
		SetRole(role).
		SetState(state).
		SetCreatedAt(now).
		SetUpdatedAt(now)

	return item
}
