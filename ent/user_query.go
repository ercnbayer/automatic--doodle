// Code generated by ent, DO NOT EDIT.

package ent

import (
	"automatic-doodle/ent/file"
	"automatic-doodle/ent/job"
	"automatic-doodle/ent/jobapplication"
	"automatic-doodle/ent/messages"
	"automatic-doodle/ent/predicate"
	"automatic-doodle/ent/refreshtoken"
	"automatic-doodle/ent/user"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	ctx                  *QueryContext
	order                []user.OrderOption
	inters               []Interceptor
	predicates           []predicate.User
	withRefreshTokens    *RefreshTokenQuery
	withProfileImage     *FileQuery
	withCoverImage       *FileQuery
	withJobs             *JobQuery
	withJobappl          *JobApplicationQuery
	withReceivedMessages *MessagesQuery
	withSentMessages     *MessagesQuery
	withFKs              bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UserQuery) Order(o ...user.OrderOption) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryRefreshTokens chains the current query on the "refresh_tokens" edge.
func (uq *UserQuery) QueryRefreshTokens() *RefreshTokenQuery {
	query := (&RefreshTokenClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(refreshtoken.Table, refreshtoken.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.RefreshTokensTable, user.RefreshTokensColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProfileImage chains the current query on the "profile_image" edge.
func (uq *UserQuery) QueryProfileImage() *FileQuery {
	query := (&FileClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, user.ProfileImageTable, user.ProfileImageColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCoverImage chains the current query on the "cover_image" edge.
func (uq *UserQuery) QueryCoverImage() *FileQuery {
	query := (&FileClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, user.CoverImageTable, user.CoverImageColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryJobs chains the current query on the "jobs" edge.
func (uq *UserQuery) QueryJobs() *JobQuery {
	query := (&JobClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(job.Table, job.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.JobsTable, user.JobsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryJobappl chains the current query on the "jobappl" edge.
func (uq *UserQuery) QueryJobappl() *JobApplicationQuery {
	query := (&JobApplicationClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(jobapplication.Table, jobapplication.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.JobapplTable, user.JobapplColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReceivedMessages chains the current query on the "received_messages" edge.
func (uq *UserQuery) QueryReceivedMessages() *MessagesQuery {
	query := (&MessagesClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(messages.Table, messages.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ReceivedMessagesTable, user.ReceivedMessagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySentMessages chains the current query on the "sent_messages" edge.
func (uq *UserQuery) QuerySentMessages() *MessagesQuery {
	query := (&MessagesClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(messages.Table, messages.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.SentMessagesTable, user.SentMessagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryAll)
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryIDs)
	if err = uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryCount)
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UserQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, ent.OpQueryExist)
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:               uq.config,
		ctx:                  uq.ctx.Clone(),
		order:                append([]user.OrderOption{}, uq.order...),
		inters:               append([]Interceptor{}, uq.inters...),
		predicates:           append([]predicate.User{}, uq.predicates...),
		withRefreshTokens:    uq.withRefreshTokens.Clone(),
		withProfileImage:     uq.withProfileImage.Clone(),
		withCoverImage:       uq.withCoverImage.Clone(),
		withJobs:             uq.withJobs.Clone(),
		withJobappl:          uq.withJobappl.Clone(),
		withReceivedMessages: uq.withReceivedMessages.Clone(),
		withSentMessages:     uq.withSentMessages.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithRefreshTokens tells the query-builder to eager-load the nodes that are connected to
// the "refresh_tokens" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithRefreshTokens(opts ...func(*RefreshTokenQuery)) *UserQuery {
	query := (&RefreshTokenClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withRefreshTokens = query
	return uq
}

// WithProfileImage tells the query-builder to eager-load the nodes that are connected to
// the "profile_image" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithProfileImage(opts ...func(*FileQuery)) *UserQuery {
	query := (&FileClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withProfileImage = query
	return uq
}

// WithCoverImage tells the query-builder to eager-load the nodes that are connected to
// the "cover_image" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithCoverImage(opts ...func(*FileQuery)) *UserQuery {
	query := (&FileClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withCoverImage = query
	return uq
}

// WithJobs tells the query-builder to eager-load the nodes that are connected to
// the "jobs" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithJobs(opts ...func(*JobQuery)) *UserQuery {
	query := (&JobClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withJobs = query
	return uq
}

// WithJobappl tells the query-builder to eager-load the nodes that are connected to
// the "jobappl" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithJobappl(opts ...func(*JobApplicationQuery)) *UserQuery {
	query := (&JobApplicationClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withJobappl = query
	return uq
}

// WithReceivedMessages tells the query-builder to eager-load the nodes that are connected to
// the "received_messages" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithReceivedMessages(opts ...func(*MessagesQuery)) *UserQuery {
	query := (&MessagesClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withReceivedMessages = query
	return uq
}

// WithSentMessages tells the query-builder to eager-load the nodes that are connected to
// the "sent_messages" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithSentMessages(opts ...func(*MessagesQuery)) *UserQuery {
	query := (&MessagesClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withSentMessages = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PhoneNumber string `json:"phone_number,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.User.Query().
//		GroupBy(user.FieldPhoneNumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = user.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PhoneNumber string `json:"phone_number,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldPhoneNumber).
//		Scan(ctx, &v)
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UserSelect{UserQuery: uq}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (uq *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		withFKs     = uq.withFKs
		_spec       = uq.querySpec()
		loadedTypes = [7]bool{
			uq.withRefreshTokens != nil,
			uq.withProfileImage != nil,
			uq.withCoverImage != nil,
			uq.withJobs != nil,
			uq.withJobappl != nil,
			uq.withReceivedMessages != nil,
			uq.withSentMessages != nil,
		}
	)
	if uq.withProfileImage != nil || uq.withCoverImage != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, user.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withRefreshTokens; query != nil {
		if err := uq.loadRefreshTokens(ctx, query, nodes,
			func(n *User) { n.Edges.RefreshTokens = []*RefreshToken{} },
			func(n *User, e *RefreshToken) { n.Edges.RefreshTokens = append(n.Edges.RefreshTokens, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withProfileImage; query != nil {
		if err := uq.loadProfileImage(ctx, query, nodes, nil,
			func(n *User, e *File) { n.Edges.ProfileImage = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withCoverImage; query != nil {
		if err := uq.loadCoverImage(ctx, query, nodes, nil,
			func(n *User, e *File) { n.Edges.CoverImage = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withJobs; query != nil {
		if err := uq.loadJobs(ctx, query, nodes,
			func(n *User) { n.Edges.Jobs = []*Job{} },
			func(n *User, e *Job) { n.Edges.Jobs = append(n.Edges.Jobs, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withJobappl; query != nil {
		if err := uq.loadJobappl(ctx, query, nodes,
			func(n *User) { n.Edges.Jobappl = []*JobApplication{} },
			func(n *User, e *JobApplication) { n.Edges.Jobappl = append(n.Edges.Jobappl, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withReceivedMessages; query != nil {
		if err := uq.loadReceivedMessages(ctx, query, nodes,
			func(n *User) { n.Edges.ReceivedMessages = []*Messages{} },
			func(n *User, e *Messages) { n.Edges.ReceivedMessages = append(n.Edges.ReceivedMessages, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withSentMessages; query != nil {
		if err := uq.loadSentMessages(ctx, query, nodes,
			func(n *User) { n.Edges.SentMessages = []*Messages{} },
			func(n *User, e *Messages) { n.Edges.SentMessages = append(n.Edges.SentMessages, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UserQuery) loadRefreshTokens(ctx context.Context, query *RefreshTokenQuery, nodes []*User, init func(*User), assign func(*User, *RefreshToken)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(refreshtoken.FieldUserID)
	}
	query.Where(predicate.RefreshToken(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.RefreshTokensColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadProfileImage(ctx context.Context, query *FileQuery, nodes []*User, init func(*User), assign func(*User, *File)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*User)
	for i := range nodes {
		if nodes[i].user_profile_image == nil {
			continue
		}
		fk := *nodes[i].user_profile_image
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(file.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_profile_image" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uq *UserQuery) loadCoverImage(ctx context.Context, query *FileQuery, nodes []*User, init func(*User), assign func(*User, *File)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*User)
	for i := range nodes {
		if nodes[i].user_cover_image == nil {
			continue
		}
		fk := *nodes[i].user_cover_image
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(file.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_cover_image" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uq *UserQuery) loadJobs(ctx context.Context, query *JobQuery, nodes []*User, init func(*User), assign func(*User, *Job)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(job.FieldJobOwner)
	}
	query.Where(predicate.Job(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.JobsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.JobOwner
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "job_owner" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadJobappl(ctx context.Context, query *JobApplicationQuery, nodes []*User, init func(*User), assign func(*User, *JobApplication)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(jobapplication.FieldUserID)
	}
	query.Where(predicate.JobApplication(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.JobapplColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadReceivedMessages(ctx context.Context, query *MessagesQuery, nodes []*User, init func(*User), assign func(*User, *Messages)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(messages.FieldTo)
	}
	query.Where(predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.ReceivedMessagesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.To
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "to" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadSentMessages(ctx context.Context, query *MessagesQuery, nodes []*User, init func(*User), assign func(*User, *Messages)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(messages.FieldFrom)
	}
	query.Where(predicate.Messages(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.SentMessagesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.From
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "from" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = user.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, ent.OpQueryGroupBy)
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, ent.OpQuerySelect)
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (us *UserSelect) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
