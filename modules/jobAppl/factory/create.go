package factory

import (
	"automatic-doodle/ent"

	"github.com/google/uuid"
)

func (f *Factory) Create(description string, pFileKey *string, jobID uuid.UUID, userID uuid.UUID) *ent.JobApplicationCreate {

	var item *ent.JobApplicationCreate

	item = f.db.JobApplication.Create().SetDescription(description).SetJobID(jobID).SetUserID(userID).SetNillableObjectKey(pFileKey)

	return item
}
