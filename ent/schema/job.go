package schema

import (
	"time"

	"entgo.io/ent"

	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Job struct {
	ent.Schema
}

func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("start_date").
			Optional(),
		field.Time("end_date").
			Optional(),
		field.Int("fee").
			Positive(),
		field.String("job_type").
			MaxLen(255).
			NotEmpty(),
		field.String("description").
			MaxLen(1024).
			NotEmpty(),
		field.UUID("job_owner", uuid.UUID{}),
	}
}

func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("jobs").
			Unique().
			Field("job_owner").
			Required(),
		edge.To("jobappl", JobApplication.Type),
	}
}

/*func (Job) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.JobFunc(func(ctx context.Context, m *ent.JobMutation) (ent.Value, error) {
					startDate, _ := m.Field("start_date")
					endDate, _ := m.Field("end_date")

					if startDate != nil && endDate != nil && endDate.(time.Time).Before(startDate.(time.Time)) {
						return nil, fmt.Errorf("end_date cannot be before start_date")
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate,
		),
	}
}*/
