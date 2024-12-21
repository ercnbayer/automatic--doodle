package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RefreshToken holds the schema definition for the RefreshToken entity.
type File struct {
	ent.Schema
}

// Fields of the RefreshToken.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.String("filename").
			MaxLen(255).
			NotEmpty(),
		field.String("extention").MaxLen(10).NotEmpty(),

		field.UUID("created_by_id", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("type").Values(
			"PROFILE_IMAGE",
			"COVER_IMAGE",
			"POST_FILE",
		),
		field.String("bucket").
			MaxLen(255).
			NotEmpty(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now), field.String("content_type").
			MaxLen(255).
			NotEmpty(),
	}
}

// Edges of the RefreshToken.
func (File) Edges() []ent.Edge {
	return nil
}
