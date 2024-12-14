package service

import "automatic-doodle/ent/file"

// TODO: this bucket names should come from env
func (svc *Service) GetBucketName(fileType file.Type) string {
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
