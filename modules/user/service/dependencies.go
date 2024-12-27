package service

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/file"
	"automatic-doodle/ent/user"
	"automatic-doodle/types"
	"context"

	"github.com/google/uuid"
)

type UserFactory interface {
	Create(string, string, string, string, string, user.Role, user.State) *ent.UserCreate
	UpdatePassword(id uuid.UUID, password string, ctx context.Context) (*ent.User, error)
}

type UserRepository interface {
	Create(*ent.UserCreate, context.Context) (*ent.User, error)
	DeleteUser(uuid.UUID, context.Context) error
	GetById(uuid.UUID, context.Context) (*ent.User, error)
	UpdateUser(ctx context.Context, u types.UpdateUserReq, id uuid.UUID) (*ent.User, error)
}

type FileRepository interface {
	Create(*ent.FileCreate, context.Context) (*ent.File, error)
}

type FileFactory interface {
	Create(string,
		string,
		string, string,
		uuid.UUID,
		file.Type) *ent.FileCreate
}
type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}
