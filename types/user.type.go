package types

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/user"

	"github.com/google/uuid"
)

type AuthenticatedUser struct {
	Id           uuid.UUID    `json:"id"`
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	Email        string       `json:"email"`
	Role         user.Role    `json:"role"`
	State        user.State   `json:"state"`
	ProfilePhoto FileResponse `json:"profile_photo"`
}

func AuthenticatedUserFromUser(u *ent.User) AuthenticatedUser {
	return AuthenticatedUser{
		Id:           u.ID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		ProfilePhoto: FromFileToFileResponse(u.Edges.ProfileImage),
		Role:         u.Role,
		State:        u.State,
	}
}
