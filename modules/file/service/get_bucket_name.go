package service

import "automatic-doodle/ent/file"

// TODO: this bucket names should come from env
func (svc *Service) GetBucketName(fileType file.Type) string {
	switch fileType {
	case file.TypePOST_FILE:
		return "post-files"
	case file.TypeCOVER_IMAGE:
		return "cover-image"
	case file.TypePROFILE_IMAGE:
		return "profile-image"
	default:
		svc.logger.Error("Undefined file type: %s", fileType)
		return "default"
	}
}
