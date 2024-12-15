package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.String("phone_number").
			MaxLen(255).
			NotEmpty().
			Unique(),
		field.String("email").
			MaxLen(255),
		field.String("first_name").MaxLen(255),
		field.String("last_name").MaxLen(255),
		field.Enum("role").Values(
			"CUSTOMER",
			"ADMIN",
		),
		field.Enum("state").Values(
			"ACTIVE",
			"DEACTIVE",
		),
		field.String("password_salt").
			NotEmpty().
			StructTag(`json:"-"`),
		field.String("password_hash").
			NotEmpty().
			StructTag(`json:"-"`),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("refresh_tokens", RefreshToken.Type),
		edge.To("profile_image", File.Type).
			Unique(),
		edge.To("cover_image", File.Type).
			Unique(),
		edge.To("jobs", Job.Type),
	}

}
