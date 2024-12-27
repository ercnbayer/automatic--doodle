package repository

import (
	"automatic-doodle/ent"
	"automatic-doodle/types"
	"context"
	"time"

	"github.com/google/uuid"
)

func (repo *UserRepository) UpdateUser(ctx context.Context, u types.UpdateUserReq, id uuid.UUID) (*ent.User, error) {
	return repo.db.User.UpdateOneID(id).SetFirstName(u.FirstName).SetLastName(u.LastName).SetEmail(u.Email).SetPhoneNumber(u.PhoneNumber).SetUpdatedAt(time.Now()).Save(ctx)
}
