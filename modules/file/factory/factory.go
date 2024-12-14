package factory

import (
	"automatic-doodle/ent"
	"automatic-doodle/ent/file"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	module     Factory
	moduleOnce sync.Once
)

func New(
	db *ent.Client,
) *Factory {
	moduleOnce.Do(func() {
		module = Factory{
			db: db,
		}
	})

	return &module
}

type Factory struct {
	db *ent.Client
}

func (f *Factory) Create(
	fileName,
	extention,
	contentType,
	bucket string,
	createdById uuid.UUID,
	fileType file.Type,
) *ent.FileCreate {
	now := time.Now()

	return f.db.File.Create().
		SetID(uuid.New()).
		SetFilename(fileName).
		SetExtention(extention).
		SetContentType(contentType).
		SetBucket(bucket).
		SetCreatedByID(createdById).
		SetType(fileType).
		SetCreatedAt(now).
		SetUpdatedAt(now)
}
