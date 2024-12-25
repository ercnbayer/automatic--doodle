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

func UserPublicDetailsFromUser(u *ent.User) UserPublicDetails {
	return UserPublicDetails{
		Id:           u.ID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		ProfilePhoto: FromFileToFileResponse(u.Edges.ProfileImage),
	}
}

type UserPublicDetails struct {
	Id           uuid.UUID    `json:"id"`
	FirstName    string       `json:"publisherfirstName"`
	LastName     string       `json:"lastName"`
	ProfilePhoto FileResponse `json:"profilePhoto"`
}
type UserPublicJobAppl struct {
	Id           uuid.UUID              `json:"id"`
	UploadedFile CreateUploadUrlRequest `json:"attachment"`
}

func UserPublicDetailsFromAUser(u AuthenticatedUser) UserPublicDetails {
	return UserPublicDetails{
		Id:           u.Id,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		ProfilePhoto: u.ProfilePhoto,
	}
}
