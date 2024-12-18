package factory

import (
	"automatic-doodle/ent"

	"github.com/google/uuid"
)

func (f *Factory) Create(description string, pFileID *uuid.UUID, jobID uuid.UUID, userID uuid.UUID) *ent.JobApplicationCreate {

	item = f.db.JobApplication.Create().SetDescription(description).SetJobID(jobID).SetUserID(userID)
}
