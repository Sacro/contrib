// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/predicate"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/user"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/withmodifiedfield"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithModifiedFieldQuery is the builder for querying WithModifiedField entities.
type WithModifiedFieldQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WithModifiedField
	// eager-loading edges.
	withOwner *UserQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WithModifiedFieldQuery builder.
func (wmfq *WithModifiedFieldQuery) Where(ps ...predicate.WithModifiedField) *WithModifiedFieldQuery {
	wmfq.predicates = append(wmfq.predicates, ps...)
	return wmfq
}

// Limit adds a limit step to the query.
func (wmfq *WithModifiedFieldQuery) Limit(limit int) *WithModifiedFieldQuery {
	wmfq.limit = &limit
	return wmfq
}

// Offset adds an offset step to the query.
func (wmfq *WithModifiedFieldQuery) Offset(offset int) *WithModifiedFieldQuery {
	wmfq.offset = &offset
	return wmfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wmfq *WithModifiedFieldQuery) Unique(unique bool) *WithModifiedFieldQuery {
	wmfq.unique = &unique
	return wmfq
}

// Order adds an order step to the query.
func (wmfq *WithModifiedFieldQuery) Order(o ...OrderFunc) *WithModifiedFieldQuery {
	wmfq.order = append(wmfq.order, o...)
	return wmfq
}

// QueryOwner chains the current query on the "owner" edge.
func (wmfq *WithModifiedFieldQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: wmfq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wmfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wmfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(withmodifiedfield.Table, withmodifiedfield.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, withmodifiedfield.OwnerTable, withmodifiedfield.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(wmfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WithModifiedField entity from the query.
// Returns a *NotFoundError when no WithModifiedField was found.
func (wmfq *WithModifiedFieldQuery) First(ctx context.Context) (*WithModifiedField, error) {
	nodes, err := wmfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{withmodifiedfield.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) FirstX(ctx context.Context) *WithModifiedField {
	node, err := wmfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WithModifiedField ID from the query.
// Returns a *NotFoundError when no WithModifiedField ID was found.
func (wmfq *WithModifiedFieldQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wmfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{withmodifiedfield.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) FirstIDX(ctx context.Context) int {
	id, err := wmfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WithModifiedField entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one WithModifiedField entity is not found.
// Returns a *NotFoundError when no WithModifiedField entities are found.
func (wmfq *WithModifiedFieldQuery) Only(ctx context.Context) (*WithModifiedField, error) {
	nodes, err := wmfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{withmodifiedfield.Label}
	default:
		return nil, &NotSingularError{withmodifiedfield.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) OnlyX(ctx context.Context) *WithModifiedField {
	node, err := wmfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WithModifiedField ID in the query.
// Returns a *NotSingularError when exactly one WithModifiedField ID is not found.
// Returns a *NotFoundError when no entities are found.
func (wmfq *WithModifiedFieldQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wmfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{withmodifiedfield.Label}
	default:
		err = &NotSingularError{withmodifiedfield.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) OnlyIDX(ctx context.Context) int {
	id, err := wmfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WithModifiedFields.
func (wmfq *WithModifiedFieldQuery) All(ctx context.Context) ([]*WithModifiedField, error) {
	if err := wmfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wmfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) AllX(ctx context.Context) []*WithModifiedField {
	nodes, err := wmfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WithModifiedField IDs.
func (wmfq *WithModifiedFieldQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wmfq.Select(withmodifiedfield.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) IDsX(ctx context.Context) []int {
	ids, err := wmfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wmfq *WithModifiedFieldQuery) Count(ctx context.Context) (int, error) {
	if err := wmfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wmfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) CountX(ctx context.Context) int {
	count, err := wmfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wmfq *WithModifiedFieldQuery) Exist(ctx context.Context) (bool, error) {
	if err := wmfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wmfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wmfq *WithModifiedFieldQuery) ExistX(ctx context.Context) bool {
	exist, err := wmfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WithModifiedFieldQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wmfq *WithModifiedFieldQuery) Clone() *WithModifiedFieldQuery {
	if wmfq == nil {
		return nil
	}
	return &WithModifiedFieldQuery{
		config:     wmfq.config,
		limit:      wmfq.limit,
		offset:     wmfq.offset,
		order:      append([]OrderFunc{}, wmfq.order...),
		predicates: append([]predicate.WithModifiedField{}, wmfq.predicates...),
		withOwner:  wmfq.withOwner.Clone(),
		// clone intermediate query.
		sql:  wmfq.sql.Clone(),
		path: wmfq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (wmfq *WithModifiedFieldQuery) WithOwner(opts ...func(*UserQuery)) *WithModifiedFieldQuery {
	query := &UserQuery{config: wmfq.config}
	for _, opt := range opts {
		opt(query)
	}
	wmfq.withOwner = query
	return wmfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WithModifiedField.Query().
//		GroupBy(withmodifiedfield.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wmfq *WithModifiedFieldQuery) GroupBy(field string, fields ...string) *WithModifiedFieldGroupBy {
	group := &WithModifiedFieldGroupBy{config: wmfq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wmfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wmfq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.WithModifiedField.Query().
//		Select(withmodifiedfield.FieldName).
//		Scan(ctx, &v)
//
func (wmfq *WithModifiedFieldQuery) Select(field string, fields ...string) *WithModifiedFieldSelect {
	wmfq.fields = append([]string{field}, fields...)
	return &WithModifiedFieldSelect{WithModifiedFieldQuery: wmfq}
}

func (wmfq *WithModifiedFieldQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wmfq.fields {
		if !withmodifiedfield.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wmfq.path != nil {
		prev, err := wmfq.path(ctx)
		if err != nil {
			return err
		}
		wmfq.sql = prev
	}
	return nil
}

func (wmfq *WithModifiedFieldQuery) sqlAll(ctx context.Context) ([]*WithModifiedField, error) {
	var (
		nodes       = []*WithModifiedField{}
		withFKs     = wmfq.withFKs
		_spec       = wmfq.querySpec()
		loadedTypes = [1]bool{
			wmfq.withOwner != nil,
		}
	)
	if wmfq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, withmodifiedfield.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &WithModifiedField{config: wmfq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, wmfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wmfq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*WithModifiedField)
		for i := range nodes {
			if nodes[i].with_modified_field_owner == nil {
				continue
			}
			fk := *nodes[i].with_modified_field_owner
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "with_modified_field_owner" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	return nodes, nil
}

func (wmfq *WithModifiedFieldQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wmfq.querySpec()
	return sqlgraph.CountNodes(ctx, wmfq.driver, _spec)
}

func (wmfq *WithModifiedFieldQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wmfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wmfq *WithModifiedFieldQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withmodifiedfield.Table,
			Columns: withmodifiedfield.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withmodifiedfield.FieldID,
			},
		},
		From:   wmfq.sql,
		Unique: true,
	}
	if unique := wmfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wmfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withmodifiedfield.FieldID)
		for i := range fields {
			if fields[i] != withmodifiedfield.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wmfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wmfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wmfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wmfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wmfq *WithModifiedFieldQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wmfq.driver.Dialect())
	t1 := builder.Table(withmodifiedfield.Table)
	selector := builder.Select(t1.Columns(withmodifiedfield.Columns...)...).From(t1)
	if wmfq.sql != nil {
		selector = wmfq.sql
		selector.Select(selector.Columns(withmodifiedfield.Columns...)...)
	}
	for _, p := range wmfq.predicates {
		p(selector)
	}
	for _, p := range wmfq.order {
		p(selector)
	}
	if offset := wmfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wmfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithModifiedFieldGroupBy is the group-by builder for WithModifiedField entities.
type WithModifiedFieldGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wmfgb *WithModifiedFieldGroupBy) Aggregate(fns ...AggregateFunc) *WithModifiedFieldGroupBy {
	wmfgb.fns = append(wmfgb.fns, fns...)
	return wmfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wmfgb *WithModifiedFieldGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wmfgb.path(ctx)
	if err != nil {
		return err
	}
	wmfgb.sql = query
	return wmfgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wmfgb *WithModifiedFieldGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wmfgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wmfgb *WithModifiedFieldGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wmfgb.fields) > 1 {
		return nil, errors.New("ent: WithModifiedFieldGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wmfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wmfgb *WithModifiedFieldGroupBy) StringsX(ctx context.Context) []string {
	v, err := wmfgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wmfgb *WithModifiedFieldGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wmfgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withmodifiedfield.Label}
	default:
		err = fmt.Errorf("ent: WithModifiedFieldGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wmfgb *WithModifiedFieldGroupBy) StringX(ctx context.Context) string {
	v, err := wmfgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wmfgb *WithModifiedFieldGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wmfgb.fields) > 1 {
		return nil, errors.New("ent: WithModifiedFieldGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wmfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wmfgb *WithModifiedFieldGroupBy) IntsX(ctx context.Context) []int {
	v, err := wmfgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wmfgb *WithModifiedFieldGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wmfgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withmodifiedfield.Label}
	default:
		err = fmt.Errorf("ent: WithModifiedFieldGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wmfgb *WithModifiedFieldGroupBy) IntX(ctx context.Context) int {
	v, err := wmfgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wmfgb *WithModifiedFieldGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wmfgb.fields) > 1 {
		return nil, errors.New("ent: WithModifiedFieldGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wmfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wmfgb *WithModifiedFieldGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wmfgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wmfgb *WithModifiedFieldGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wmfgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withmodifiedfield.Label}
	default:
		err = fmt.Errorf("ent: WithModifiedFieldGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wmfgb *WithModifiedFieldGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wmfgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wmfgb *WithModifiedFieldGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wmfgb.fields) > 1 {
		return nil, errors.New("ent: WithModifiedFieldGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wmfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wmfgb *WithModifiedFieldGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wmfgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wmfgb *WithModifiedFieldGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wmfgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withmodifiedfield.Label}
	default:
		err = fmt.Errorf("ent: WithModifiedFieldGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wmfgb *WithModifiedFieldGroupBy) BoolX(ctx context.Context) bool {
	v, err := wmfgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wmfgb *WithModifiedFieldGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wmfgb.fields {
		if !withmodifiedfield.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wmfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wmfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wmfgb *WithModifiedFieldGroupBy) sqlQuery() *sql.Selector {
	selector := wmfgb.sql
	columns := make([]string, 0, len(wmfgb.fields)+len(wmfgb.fns))
	columns = append(columns, wmfgb.fields...)
	for _, fn := range wmfgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(wmfgb.fields...)
}

// WithModifiedFieldSelect is the builder for selecting fields of WithModifiedField entities.
type WithModifiedFieldSelect struct {
	*WithModifiedFieldQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wmfs *WithModifiedFieldSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wmfs.prepareQuery(ctx); err != nil {
		return err
	}
	wmfs.sql = wmfs.WithModifiedFieldQuery.sqlQuery(ctx)
	return wmfs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wmfs *WithModifiedFieldSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wmfs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (wmfs *WithModifiedFieldSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wmfs.fields) > 1 {
		return nil, errors.New("ent: WithModifiedFieldSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wmfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wmfs *WithModifiedFieldSelect) StringsX(ctx context.Context) []string {
	v, err := wmfs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (wmfs *WithModifiedFieldSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wmfs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withmodifiedfield.Label}
	default:
		err = fmt.Errorf("ent: WithModifiedFieldSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wmfs *WithModifiedFieldSelect) StringX(ctx context.Context) string {
	v, err := wmfs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (wmfs *WithModifiedFieldSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wmfs.fields) > 1 {
		return nil, errors.New("ent: WithModifiedFieldSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wmfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wmfs *WithModifiedFieldSelect) IntsX(ctx context.Context) []int {
	v, err := wmfs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (wmfs *WithModifiedFieldSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wmfs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withmodifiedfield.Label}
	default:
		err = fmt.Errorf("ent: WithModifiedFieldSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wmfs *WithModifiedFieldSelect) IntX(ctx context.Context) int {
	v, err := wmfs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (wmfs *WithModifiedFieldSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wmfs.fields) > 1 {
		return nil, errors.New("ent: WithModifiedFieldSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wmfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wmfs *WithModifiedFieldSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wmfs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (wmfs *WithModifiedFieldSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wmfs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withmodifiedfield.Label}
	default:
		err = fmt.Errorf("ent: WithModifiedFieldSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wmfs *WithModifiedFieldSelect) Float64X(ctx context.Context) float64 {
	v, err := wmfs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (wmfs *WithModifiedFieldSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wmfs.fields) > 1 {
		return nil, errors.New("ent: WithModifiedFieldSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wmfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wmfs *WithModifiedFieldSelect) BoolsX(ctx context.Context) []bool {
	v, err := wmfs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (wmfs *WithModifiedFieldSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wmfs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withmodifiedfield.Label}
	default:
		err = fmt.Errorf("ent: WithModifiedFieldSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wmfs *WithModifiedFieldSelect) BoolX(ctx context.Context) bool {
	v, err := wmfs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wmfs *WithModifiedFieldSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wmfs.sqlQuery().Query()
	if err := wmfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wmfs *WithModifiedFieldSelect) sqlQuery() sql.Querier {
	selector := wmfs.sql
	selector.Select(selector.Columns(wmfs.fields...)...)
	return selector
}
