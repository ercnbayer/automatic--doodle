package service

import (
	"automatic-doodle/ent/file"
	"automatic-doodle/types"
	"context"
	"mime"
	"path/filepath"

	"github.com/google/uuid"
)

func (svc *UserService) GetBucketName(fileType file.Type) string {
	switch fileType {
	case file.TypePOST_FILE:
		return "post_files"
	case file.TypeCOVER_IMAGE:
		return "cover_image"
	case file.TypePROFILE_IMAGE:
		return "profile_image"
	default:
		svc.logger.Error("Undefined file type: %s", fileType)
		return "default"
	}
}

func (svc *UserService) UpdateProfilePhoto(
	userId uuid.UUID,
	request *types.UpdateProfilePhotoRequest,
) (types.AuthenticatedUser, error) {
	fileExt := filepath.Ext(request.FileKey)
	mimeType := mime.TypeByExtension(fileExt)
	fileCreate := svc.fileFactory.Create(
		request.FileKey,
		fileExt,
		mimeType,
		svc.GetBucketName(file.TypePROFILE_IMAGE),
		userId,
		file.TypePROFILE_IMAGE,
	)
	_, err := svc.fileRepository.Create(fileCreate, context.Background())
	if err != nil {
		svc.logger.Error(
			`Something went wrong during UpdateProfilePhoto->FileCreate: %w`,
			err,
		)
		return types.AuthenticatedUser{}, err
	}

	userItem, err := svc.userRepository.GetById(userId)
	if err != nil {
		svc.logger.Error(
			`Something went wrong during UpdateProfilePhoto->GetUserById: %w`,
			err,
		)
		return types.AuthenticatedUser{}, err
	}

	return types.AuthenticatedUserFromUser(userItem), nil
}
