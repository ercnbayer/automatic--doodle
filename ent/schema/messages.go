package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Messages holds the schema definition for the Messages entity.
type Messages struct {
	ent.Schema
}

// Fields of the Messages.
func (Messages) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.UUID("from", uuid.UUID{}),

		field.UUID("to", uuid.UUID{}),

		field.String("message").NotEmpty(),

		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Messages.
func (Messages) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", User.Type).
			Ref("sent_messages").
			Field("from").
			Unique().
			Required(),

		edge.From("receiver", User.Type).
			Ref("received_messages").
			Field("to").
			Unique().
			Required(),
	}
}
