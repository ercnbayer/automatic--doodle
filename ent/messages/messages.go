// Code generated by ent, DO NOT EDIT.

package messages

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the messages type in the database.
	Label = "messages"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFrom holds the string denoting the from field in the database.
	FieldFrom = "from"
	// FieldTo holds the string denoting the to field in the database.
	FieldTo = "to"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeSender holds the string denoting the sender edge name in mutations.
	EdgeSender = "sender"
	// EdgeReceiver holds the string denoting the receiver edge name in mutations.
	EdgeReceiver = "receiver"
	// Table holds the table name of the messages in the database.
	Table = "messages"
	// SenderTable is the table that holds the sender relation/edge.
	SenderTable = "messages"
	// SenderInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SenderInverseTable = "users"
	// SenderColumn is the table column denoting the sender relation/edge.
	SenderColumn = "from"
	// ReceiverTable is the table that holds the receiver relation/edge.
	ReceiverTable = "messages"
	// ReceiverInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ReceiverInverseTable = "users"
	// ReceiverColumn is the table column denoting the receiver relation/edge.
	ReceiverColumn = "to"
)

// Columns holds all SQL columns for messages fields.
var Columns = []string{
	FieldID,
	FieldFrom,
	FieldTo,
	FieldMessage,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// MessageValidator is a validator for the "message" field. It is called by the builders before save.
	MessageValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Messages queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFrom orders the results by the from field.
func ByFrom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFrom, opts...).ToFunc()
}

// ByTo orders the results by the to field.
func ByTo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTo, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// BySenderField orders the results by sender field.
func BySenderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSenderStep(), sql.OrderByField(field, opts...))
	}
}

// ByReceiverField orders the results by receiver field.
func ByReceiverField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReceiverStep(), sql.OrderByField(field, opts...))
	}
}
func newSenderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SenderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SenderTable, SenderColumn),
	)
}
func newReceiverStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReceiverInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ReceiverTable, ReceiverColumn),
	)
}
