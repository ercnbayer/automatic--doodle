// Code generated by ent, DO NOT EDIT.

package ent

import (
	"automatic-doodle/ent/file"
	"automatic-doodle/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Role holds the value of the "role" field.
	Role user.Role `json:"role,omitempty"`
	// State holds the value of the "state" field.
	State user.State `json:"state,omitempty"`
	// PasswordSalt holds the value of the "password_salt" field.
	PasswordSalt string `json:"-"`
	// PasswordHash holds the value of the "password_hash" field.
	PasswordHash string `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges              UserEdges `json:"edges"`
	user_profile_image *uuid.UUID
	user_cover_image   *uuid.UUID
	selectValues       sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// RefreshTokens holds the value of the refresh_tokens edge.
	RefreshTokens []*RefreshToken `json:"refresh_tokens,omitempty"`
	// ProfileImage holds the value of the profile_image edge.
	ProfileImage *File `json:"profile_image,omitempty"`
	// CoverImage holds the value of the cover_image edge.
	CoverImage *File `json:"cover_image,omitempty"`
	// Jobs holds the value of the jobs edge.
	Jobs []*Job `json:"jobs,omitempty"`
	// Jobappl holds the value of the jobappl edge.
	Jobappl []*JobApplication `json:"jobappl,omitempty"`
	// ReceivedMessages holds the value of the received_messages edge.
	ReceivedMessages []*Messages `json:"received_messages,omitempty"`
	// SentMessages holds the value of the sent_messages edge.
	SentMessages []*Messages `json:"sent_messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// RefreshTokensOrErr returns the RefreshTokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RefreshTokensOrErr() ([]*RefreshToken, error) {
	if e.loadedTypes[0] {
		return e.RefreshTokens, nil
	}
	return nil, &NotLoadedError{edge: "refresh_tokens"}
}

// ProfileImageOrErr returns the ProfileImage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ProfileImageOrErr() (*File, error) {
	if e.ProfileImage != nil {
		return e.ProfileImage, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: file.Label}
	}
	return nil, &NotLoadedError{edge: "profile_image"}
}

// CoverImageOrErr returns the CoverImage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CoverImageOrErr() (*File, error) {
	if e.CoverImage != nil {
		return e.CoverImage, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: file.Label}
	}
	return nil, &NotLoadedError{edge: "cover_image"}
}

// JobsOrErr returns the Jobs value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) JobsOrErr() ([]*Job, error) {
	if e.loadedTypes[3] {
		return e.Jobs, nil
	}
	return nil, &NotLoadedError{edge: "jobs"}
}

// JobapplOrErr returns the Jobappl value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) JobapplOrErr() ([]*JobApplication, error) {
	if e.loadedTypes[4] {
		return e.Jobappl, nil
	}
	return nil, &NotLoadedError{edge: "jobappl"}
}

// ReceivedMessagesOrErr returns the ReceivedMessages value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReceivedMessagesOrErr() ([]*Messages, error) {
	if e.loadedTypes[5] {
		return e.ReceivedMessages, nil
	}
	return nil, &NotLoadedError{edge: "received_messages"}
}

// SentMessagesOrErr returns the SentMessages value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SentMessagesOrErr() ([]*Messages, error) {
	if e.loadedTypes[6] {
		return e.SentMessages, nil
	}
	return nil, &NotLoadedError{edge: "sent_messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldPhoneNumber, user.FieldEmail, user.FieldFirstName, user.FieldLastName, user.FieldRole, user.FieldState, user.FieldPasswordSalt, user.FieldPasswordHash:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		case user.ForeignKeys[0]: // user_profile_image
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case user.ForeignKeys[1]: // user_cover_image
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				u.PhoneNumber = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = user.Role(value.String)
			}
		case user.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				u.State = user.State(value.String)
			}
		case user.FieldPasswordSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_salt", values[i])
			} else if value.Valid {
				u.PasswordSalt = value.String
			}
		case user.FieldPasswordHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password_hash", values[i])
			} else if value.Valid {
				u.PasswordHash = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_profile_image", values[i])
			} else if value.Valid {
				u.user_profile_image = new(uuid.UUID)
				*u.user_profile_image = *value.S.(*uuid.UUID)
			}
		case user.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_cover_image", values[i])
			} else if value.Valid {
				u.user_cover_image = new(uuid.UUID)
				*u.user_cover_image = *value.S.(*uuid.UUID)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryRefreshTokens queries the "refresh_tokens" edge of the User entity.
func (u *User) QueryRefreshTokens() *RefreshTokenQuery {
	return NewUserClient(u.config).QueryRefreshTokens(u)
}

// QueryProfileImage queries the "profile_image" edge of the User entity.
func (u *User) QueryProfileImage() *FileQuery {
	return NewUserClient(u.config).QueryProfileImage(u)
}

// QueryCoverImage queries the "cover_image" edge of the User entity.
func (u *User) QueryCoverImage() *FileQuery {
	return NewUserClient(u.config).QueryCoverImage(u)
}

// QueryJobs queries the "jobs" edge of the User entity.
func (u *User) QueryJobs() *JobQuery {
	return NewUserClient(u.config).QueryJobs(u)
}

// QueryJobappl queries the "jobappl" edge of the User entity.
func (u *User) QueryJobappl() *JobApplicationQuery {
	return NewUserClient(u.config).QueryJobappl(u)
}

// QueryReceivedMessages queries the "received_messages" edge of the User entity.
func (u *User) QueryReceivedMessages() *MessagesQuery {
	return NewUserClient(u.config).QueryReceivedMessages(u)
}

// QuerySentMessages queries the "sent_messages" edge of the User entity.
func (u *User) QuerySentMessages() *MessagesQuery {
	return NewUserClient(u.config).QuerySentMessages(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("phone_number=")
	builder.WriteString(u.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", u.State))
	builder.WriteString(", ")
	builder.WriteString("password_salt=")
	builder.WriteString(u.PasswordSalt)
	builder.WriteString(", ")
	builder.WriteString("password_hash=")
	builder.WriteString(u.PasswordHash)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
