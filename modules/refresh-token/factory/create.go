package factory

import (
	"automatic-doodle/ent"
	"time"

	"github.com/google/uuid"
)

func (f *Factory) Create(
	userId uuid.UUID,
) *ent.RefreshTokenCreate {
	now := time.Now()
	token := f.encryption.GenerateSalt(16)

	return f.db.RefreshToken.Create().
		SetID(uuid.New()).
		SetToken(token).
		SetIsClaimed(false).
		SetUserID(userId).
		SetExpiredAt(time.Now().Add(time.Hour * 1)).
		SetCreatedAt(now).
		SetUpdatedAt(now)
}
