package types

import (
	"automatic-doodle/ent/user"

	"github.com/google/uuid"
)

type AuthenticatedUser struct {
	Id        uuid.UUID  `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Role      user.Role  `json:"role"`
	State     user.State `json:"state"`
}
