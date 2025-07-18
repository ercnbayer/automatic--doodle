// Code generated by ent, DO NOT EDIT.

package jobapplication

import (
	"automatic-doodle/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldLTE(FieldID, id))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEQ(FieldDescription, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEQ(FieldUserID, v))
}

// JobID applies equality check predicate on the "job_id" field. It's identical to JobIDEQ.
func JobID(v uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEQ(FieldJobID, v))
}

// ObjectKey applies equality check predicate on the "object_key" field. It's identical to ObjectKeyEQ.
func ObjectKey(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEQ(FieldObjectKey, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldContainsFold(FieldDescription, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNotIn(FieldUserID, vs...))
}

// JobIDEQ applies the EQ predicate on the "job_id" field.
func JobIDEQ(v uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEQ(FieldJobID, v))
}

// JobIDNEQ applies the NEQ predicate on the "job_id" field.
func JobIDNEQ(v uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNEQ(FieldJobID, v))
}

// JobIDIn applies the In predicate on the "job_id" field.
func JobIDIn(vs ...uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldIn(FieldJobID, vs...))
}

// JobIDNotIn applies the NotIn predicate on the "job_id" field.
func JobIDNotIn(vs ...uuid.UUID) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNotIn(FieldJobID, vs...))
}

// ObjectKeyEQ applies the EQ predicate on the "object_key" field.
func ObjectKeyEQ(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEQ(FieldObjectKey, v))
}

// ObjectKeyNEQ applies the NEQ predicate on the "object_key" field.
func ObjectKeyNEQ(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNEQ(FieldObjectKey, v))
}

// ObjectKeyIn applies the In predicate on the "object_key" field.
func ObjectKeyIn(vs ...string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldIn(FieldObjectKey, vs...))
}

// ObjectKeyNotIn applies the NotIn predicate on the "object_key" field.
func ObjectKeyNotIn(vs ...string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNotIn(FieldObjectKey, vs...))
}

// ObjectKeyGT applies the GT predicate on the "object_key" field.
func ObjectKeyGT(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldGT(FieldObjectKey, v))
}

// ObjectKeyGTE applies the GTE predicate on the "object_key" field.
func ObjectKeyGTE(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldGTE(FieldObjectKey, v))
}

// ObjectKeyLT applies the LT predicate on the "object_key" field.
func ObjectKeyLT(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldLT(FieldObjectKey, v))
}

// ObjectKeyLTE applies the LTE predicate on the "object_key" field.
func ObjectKeyLTE(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldLTE(FieldObjectKey, v))
}

// ObjectKeyContains applies the Contains predicate on the "object_key" field.
func ObjectKeyContains(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldContains(FieldObjectKey, v))
}

// ObjectKeyHasPrefix applies the HasPrefix predicate on the "object_key" field.
func ObjectKeyHasPrefix(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldHasPrefix(FieldObjectKey, v))
}

// ObjectKeyHasSuffix applies the HasSuffix predicate on the "object_key" field.
func ObjectKeyHasSuffix(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldHasSuffix(FieldObjectKey, v))
}

// ObjectKeyIsNil applies the IsNil predicate on the "object_key" field.
func ObjectKeyIsNil() predicate.JobApplication {
	return predicate.JobApplication(sql.FieldIsNull(FieldObjectKey))
}

// ObjectKeyNotNil applies the NotNil predicate on the "object_key" field.
func ObjectKeyNotNil() predicate.JobApplication {
	return predicate.JobApplication(sql.FieldNotNull(FieldObjectKey))
}

// ObjectKeyEqualFold applies the EqualFold predicate on the "object_key" field.
func ObjectKeyEqualFold(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldEqualFold(FieldObjectKey, v))
}

// ObjectKeyContainsFold applies the ContainsFold predicate on the "object_key" field.
func ObjectKeyContainsFold(v string) predicate.JobApplication {
	return predicate.JobApplication(sql.FieldContainsFold(FieldObjectKey, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.JobApplication {
	return predicate.JobApplication(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.JobApplication {
	return predicate.JobApplication(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasJob applies the HasEdge predicate on the "job" edge.
func HasJob() predicate.JobApplication {
	return predicate.JobApplication(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, JobTable, JobColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJobWith applies the HasEdge predicate on the "job" edge with a given conditions (other predicates).
func HasJobWith(preds ...predicate.Job) predicate.JobApplication {
	return predicate.JobApplication(func(s *sql.Selector) {
		step := newJobStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobApplication) predicate.JobApplication {
	return predicate.JobApplication(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobApplication) predicate.JobApplication {
	return predicate.JobApplication(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.JobApplication) predicate.JobApplication {
	return predicate.JobApplication(sql.NotPredicates(p))
}
