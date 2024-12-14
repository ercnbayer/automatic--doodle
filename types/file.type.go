package types

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/file"
)

type CreateUploadUrlRequest struct {
	Filename string    `json:"fileName" validate:"required,filename,fileext"`
	Type     file.Type `json:"type" validate:"required,oneof=PROFILE_IMAGE COVER_IMAGE POST_FILE"`
}

type CreateUploadUrlResponse struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

type FileResponse struct {
	Key    string `json:"key"`
	Bucket string `json:"bucket"`
}

func FromFileToFileResponse(item *ent.File) FileResponse {
	if item == nil {
		return FileResponse{}
	}

	return FileResponse{
		Key:    item.Filename,
		Bucket: item.Bucket,
	}
}
