package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type JobApplication struct {
	ent.Schema
}

func (JobApplication) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.String("description").
			MaxLen(1024),
		field.UUID("user_id", uuid.UUID{}), // Foreign key to User

		field.UUID("job_id", uuid.UUID{}), // Foreign key to Job

		field.UUID("file_id", uuid.UUID{}), // Foreign key to File

	}
}

func (JobApplication) Edges() []ent.Edge {
	return []ent.Edge{
		// User edge: connecting job application with user
		edge.From("user", User.Type).
			Ref("jobappl").
			Field("user_id"). // Foreign key
			Required().Unique(),

		// Job edge: connecting job application with job
		edge.From("job", Job.Type).
			Ref("jobappl").
			Field("job_id"). // Foreign key
			Required().Unique(),

		// File edge: connecting job application with file
		edge.From("file", File.Type).
			Ref("jobappl").
			Field("file_id"). // Foreign key
			Required().Unique(),
	}
}
