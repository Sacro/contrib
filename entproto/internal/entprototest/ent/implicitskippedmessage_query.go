// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImplicitSkippedMessageQuery is the builder for querying ImplicitSkippedMessage entities.
type ImplicitSkippedMessageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ImplicitSkippedMessage
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImplicitSkippedMessageQuery builder.
func (ismq *ImplicitSkippedMessageQuery) Where(ps ...predicate.ImplicitSkippedMessage) *ImplicitSkippedMessageQuery {
	ismq.predicates = append(ismq.predicates, ps...)
	return ismq
}

// Limit adds a limit step to the query.
func (ismq *ImplicitSkippedMessageQuery) Limit(limit int) *ImplicitSkippedMessageQuery {
	ismq.limit = &limit
	return ismq
}

// Offset adds an offset step to the query.
func (ismq *ImplicitSkippedMessageQuery) Offset(offset int) *ImplicitSkippedMessageQuery {
	ismq.offset = &offset
	return ismq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ismq *ImplicitSkippedMessageQuery) Unique(unique bool) *ImplicitSkippedMessageQuery {
	ismq.unique = &unique
	return ismq
}

// Order adds an order step to the query.
func (ismq *ImplicitSkippedMessageQuery) Order(o ...OrderFunc) *ImplicitSkippedMessageQuery {
	ismq.order = append(ismq.order, o...)
	return ismq
}

// First returns the first ImplicitSkippedMessage entity from the query.
// Returns a *NotFoundError when no ImplicitSkippedMessage was found.
func (ismq *ImplicitSkippedMessageQuery) First(ctx context.Context) (*ImplicitSkippedMessage, error) {
	nodes, err := ismq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{implicitskippedmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) FirstX(ctx context.Context) *ImplicitSkippedMessage {
	node, err := ismq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ImplicitSkippedMessage ID from the query.
// Returns a *NotFoundError when no ImplicitSkippedMessage ID was found.
func (ismq *ImplicitSkippedMessageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ismq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{implicitskippedmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) FirstIDX(ctx context.Context) int {
	id, err := ismq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ImplicitSkippedMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ImplicitSkippedMessage entity is not found.
// Returns a *NotFoundError when no ImplicitSkippedMessage entities are found.
func (ismq *ImplicitSkippedMessageQuery) Only(ctx context.Context) (*ImplicitSkippedMessage, error) {
	nodes, err := ismq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{implicitskippedmessage.Label}
	default:
		return nil, &NotSingularError{implicitskippedmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) OnlyX(ctx context.Context) *ImplicitSkippedMessage {
	node, err := ismq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ImplicitSkippedMessage ID in the query.
// Returns a *NotSingularError when exactly one ImplicitSkippedMessage ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ismq *ImplicitSkippedMessageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ismq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = &NotSingularError{implicitskippedmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) OnlyIDX(ctx context.Context) int {
	id, err := ismq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ImplicitSkippedMessages.
func (ismq *ImplicitSkippedMessageQuery) All(ctx context.Context) ([]*ImplicitSkippedMessage, error) {
	if err := ismq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ismq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) AllX(ctx context.Context) []*ImplicitSkippedMessage {
	nodes, err := ismq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ImplicitSkippedMessage IDs.
func (ismq *ImplicitSkippedMessageQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ismq.Select(implicitskippedmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) IDsX(ctx context.Context) []int {
	ids, err := ismq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ismq *ImplicitSkippedMessageQuery) Count(ctx context.Context) (int, error) {
	if err := ismq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ismq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) CountX(ctx context.Context) int {
	count, err := ismq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ismq *ImplicitSkippedMessageQuery) Exist(ctx context.Context) (bool, error) {
	if err := ismq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ismq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := ismq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImplicitSkippedMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ismq *ImplicitSkippedMessageQuery) Clone() *ImplicitSkippedMessageQuery {
	if ismq == nil {
		return nil
	}
	return &ImplicitSkippedMessageQuery{
		config:     ismq.config,
		limit:      ismq.limit,
		offset:     ismq.offset,
		order:      append([]OrderFunc{}, ismq.order...),
		predicates: append([]predicate.ImplicitSkippedMessage{}, ismq.predicates...),
		// clone intermediate query.
		sql:  ismq.sql.Clone(),
		path: ismq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (ismq *ImplicitSkippedMessageQuery) GroupBy(field string, fields ...string) *ImplicitSkippedMessageGroupBy {
	group := &ImplicitSkippedMessageGroupBy{config: ismq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ismq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ismq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (ismq *ImplicitSkippedMessageQuery) Select(field string, fields ...string) *ImplicitSkippedMessageSelect {
	ismq.fields = append([]string{field}, fields...)
	return &ImplicitSkippedMessageSelect{ImplicitSkippedMessageQuery: ismq}
}

func (ismq *ImplicitSkippedMessageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ismq.fields {
		if !implicitskippedmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ismq.path != nil {
		prev, err := ismq.path(ctx)
		if err != nil {
			return err
		}
		ismq.sql = prev
	}
	return nil
}

func (ismq *ImplicitSkippedMessageQuery) sqlAll(ctx context.Context) ([]*ImplicitSkippedMessage, error) {
	var (
		nodes   = []*ImplicitSkippedMessage{}
		withFKs = ismq.withFKs
		_spec   = ismq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, implicitskippedmessage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ImplicitSkippedMessage{config: ismq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ismq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ismq *ImplicitSkippedMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ismq.querySpec()
	return sqlgraph.CountNodes(ctx, ismq.driver, _spec)
}

func (ismq *ImplicitSkippedMessageQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ismq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ismq *ImplicitSkippedMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   implicitskippedmessage.Table,
			Columns: implicitskippedmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: implicitskippedmessage.FieldID,
			},
		},
		From:   ismq.sql,
		Unique: true,
	}
	if unique := ismq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ismq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, implicitskippedmessage.FieldID)
		for i := range fields {
			if fields[i] != implicitskippedmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ismq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ismq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ismq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ismq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ismq *ImplicitSkippedMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ismq.driver.Dialect())
	t1 := builder.Table(implicitskippedmessage.Table)
	selector := builder.Select(t1.Columns(implicitskippedmessage.Columns...)...).From(t1)
	if ismq.sql != nil {
		selector = ismq.sql
		selector.Select(selector.Columns(implicitskippedmessage.Columns...)...)
	}
	for _, p := range ismq.predicates {
		p(selector)
	}
	for _, p := range ismq.order {
		p(selector)
	}
	if offset := ismq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ismq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImplicitSkippedMessageGroupBy is the group-by builder for ImplicitSkippedMessage entities.
type ImplicitSkippedMessageGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ismgb *ImplicitSkippedMessageGroupBy) Aggregate(fns ...AggregateFunc) *ImplicitSkippedMessageGroupBy {
	ismgb.fns = append(ismgb.fns, fns...)
	return ismgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ismgb *ImplicitSkippedMessageGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ismgb.path(ctx)
	if err != nil {
		return err
	}
	ismgb.sql = query
	return ismgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ismgb *ImplicitSkippedMessageGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ismgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ismgb *ImplicitSkippedMessageGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ismgb.fields) > 1 {
		return nil, errors.New("ent: ImplicitSkippedMessageGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ismgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ismgb *ImplicitSkippedMessageGroupBy) StringsX(ctx context.Context) []string {
	v, err := ismgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ismgb *ImplicitSkippedMessageGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ismgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = fmt.Errorf("ent: ImplicitSkippedMessageGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ismgb *ImplicitSkippedMessageGroupBy) StringX(ctx context.Context) string {
	v, err := ismgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ismgb *ImplicitSkippedMessageGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ismgb.fields) > 1 {
		return nil, errors.New("ent: ImplicitSkippedMessageGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ismgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ismgb *ImplicitSkippedMessageGroupBy) IntsX(ctx context.Context) []int {
	v, err := ismgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ismgb *ImplicitSkippedMessageGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ismgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = fmt.Errorf("ent: ImplicitSkippedMessageGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ismgb *ImplicitSkippedMessageGroupBy) IntX(ctx context.Context) int {
	v, err := ismgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ismgb *ImplicitSkippedMessageGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ismgb.fields) > 1 {
		return nil, errors.New("ent: ImplicitSkippedMessageGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ismgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ismgb *ImplicitSkippedMessageGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ismgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ismgb *ImplicitSkippedMessageGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ismgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = fmt.Errorf("ent: ImplicitSkippedMessageGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ismgb *ImplicitSkippedMessageGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ismgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ismgb *ImplicitSkippedMessageGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ismgb.fields) > 1 {
		return nil, errors.New("ent: ImplicitSkippedMessageGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ismgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ismgb *ImplicitSkippedMessageGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ismgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ismgb *ImplicitSkippedMessageGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ismgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = fmt.Errorf("ent: ImplicitSkippedMessageGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ismgb *ImplicitSkippedMessageGroupBy) BoolX(ctx context.Context) bool {
	v, err := ismgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ismgb *ImplicitSkippedMessageGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ismgb.fields {
		if !implicitskippedmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ismgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ismgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ismgb *ImplicitSkippedMessageGroupBy) sqlQuery() *sql.Selector {
	selector := ismgb.sql
	columns := make([]string, 0, len(ismgb.fields)+len(ismgb.fns))
	columns = append(columns, ismgb.fields...)
	for _, fn := range ismgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ismgb.fields...)
}

// ImplicitSkippedMessageSelect is the builder for selecting fields of ImplicitSkippedMessage entities.
type ImplicitSkippedMessageSelect struct {
	*ImplicitSkippedMessageQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (isms *ImplicitSkippedMessageSelect) Scan(ctx context.Context, v interface{}) error {
	if err := isms.prepareQuery(ctx); err != nil {
		return err
	}
	isms.sql = isms.ImplicitSkippedMessageQuery.sqlQuery(ctx)
	return isms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (isms *ImplicitSkippedMessageSelect) ScanX(ctx context.Context, v interface{}) {
	if err := isms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (isms *ImplicitSkippedMessageSelect) Strings(ctx context.Context) ([]string, error) {
	if len(isms.fields) > 1 {
		return nil, errors.New("ent: ImplicitSkippedMessageSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := isms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (isms *ImplicitSkippedMessageSelect) StringsX(ctx context.Context) []string {
	v, err := isms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (isms *ImplicitSkippedMessageSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = isms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = fmt.Errorf("ent: ImplicitSkippedMessageSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (isms *ImplicitSkippedMessageSelect) StringX(ctx context.Context) string {
	v, err := isms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (isms *ImplicitSkippedMessageSelect) Ints(ctx context.Context) ([]int, error) {
	if len(isms.fields) > 1 {
		return nil, errors.New("ent: ImplicitSkippedMessageSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := isms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (isms *ImplicitSkippedMessageSelect) IntsX(ctx context.Context) []int {
	v, err := isms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (isms *ImplicitSkippedMessageSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = isms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = fmt.Errorf("ent: ImplicitSkippedMessageSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (isms *ImplicitSkippedMessageSelect) IntX(ctx context.Context) int {
	v, err := isms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (isms *ImplicitSkippedMessageSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(isms.fields) > 1 {
		return nil, errors.New("ent: ImplicitSkippedMessageSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := isms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (isms *ImplicitSkippedMessageSelect) Float64sX(ctx context.Context) []float64 {
	v, err := isms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (isms *ImplicitSkippedMessageSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = isms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = fmt.Errorf("ent: ImplicitSkippedMessageSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (isms *ImplicitSkippedMessageSelect) Float64X(ctx context.Context) float64 {
	v, err := isms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (isms *ImplicitSkippedMessageSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(isms.fields) > 1 {
		return nil, errors.New("ent: ImplicitSkippedMessageSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := isms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (isms *ImplicitSkippedMessageSelect) BoolsX(ctx context.Context) []bool {
	v, err := isms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (isms *ImplicitSkippedMessageSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = isms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = fmt.Errorf("ent: ImplicitSkippedMessageSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (isms *ImplicitSkippedMessageSelect) BoolX(ctx context.Context) bool {
	v, err := isms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (isms *ImplicitSkippedMessageSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := isms.sqlQuery().Query()
	if err := isms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (isms *ImplicitSkippedMessageSelect) sqlQuery() sql.Querier {
	selector := isms.sql
	selector.Select(selector.Columns(isms.fields...)...)
	return selector
}
