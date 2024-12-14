package types

type UpdateProfilePhotoRequest struct {
	FileKey string `json:"fileKey" validate:"required,filename,fileext"`
}
