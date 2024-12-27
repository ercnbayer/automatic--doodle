package factory

import (
	"automatic-doodle/ent"
	"context"
	"time"

	"github.com/google/uuid"
)

func (f *UserFactory) UpdatePassword(id uuid.UUID, password string, ctx context.Context) (*ent.User, error) {

	salt := f.encryption.GenerateSalt(16)

	hash := f.encryption.HashPassword(password, salt)

	return f.db.User.UpdateOneID(id).SetPasswordHash(hash).SetPasswordSalt(salt).SetUpdatedAt(time.Now()).Save(ctx)
}
