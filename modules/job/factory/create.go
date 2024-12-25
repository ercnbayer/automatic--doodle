package factory

import (
	"automatic-doodle/ent"
	"time"

	"github.com/google/uuid"
)

func (f *Factory) Create(fee int, desc string, jobType string, startDate time.Time, endDate time.Time, publisherID uuid.UUID) *ent.JobCreate {
	return f.db.Job.Create().SetFee(fee).SetDescription(desc).SetJobType(jobType).SetStartDate(startDate).SetEndDate(endDate).SetJobOwner(publisherID)
}
