// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/predicate"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/withoutfields"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithoutFieldsQuery is the builder for querying WithoutFields entities.
type WithoutFieldsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WithoutFields
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WithoutFieldsQuery builder.
func (wfq *WithoutFieldsQuery) Where(ps ...predicate.WithoutFields) *WithoutFieldsQuery {
	wfq.predicates = append(wfq.predicates, ps...)
	return wfq
}

// Limit adds a limit step to the query.
func (wfq *WithoutFieldsQuery) Limit(limit int) *WithoutFieldsQuery {
	wfq.limit = &limit
	return wfq
}

// Offset adds an offset step to the query.
func (wfq *WithoutFieldsQuery) Offset(offset int) *WithoutFieldsQuery {
	wfq.offset = &offset
	return wfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wfq *WithoutFieldsQuery) Unique(unique bool) *WithoutFieldsQuery {
	wfq.unique = &unique
	return wfq
}

// Order adds an order step to the query.
func (wfq *WithoutFieldsQuery) Order(o ...OrderFunc) *WithoutFieldsQuery {
	wfq.order = append(wfq.order, o...)
	return wfq
}

// First returns the first WithoutFields entity from the query.
// Returns a *NotFoundError when no WithoutFields was found.
func (wfq *WithoutFieldsQuery) First(ctx context.Context) (*WithoutFields, error) {
	nodes, err := wfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{withoutfields.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) FirstX(ctx context.Context) *WithoutFields {
	node, err := wfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WithoutFields ID from the query.
// Returns a *NotFoundError when no WithoutFields ID was found.
func (wfq *WithoutFieldsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{withoutfields.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) FirstIDX(ctx context.Context) int {
	id, err := wfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WithoutFields entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one WithoutFields entity is not found.
// Returns a *NotFoundError when no WithoutFields entities are found.
func (wfq *WithoutFieldsQuery) Only(ctx context.Context) (*WithoutFields, error) {
	nodes, err := wfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{withoutfields.Label}
	default:
		return nil, &NotSingularError{withoutfields.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) OnlyX(ctx context.Context) *WithoutFields {
	node, err := wfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WithoutFields ID in the query.
// Returns a *NotSingularError when exactly one WithoutFields ID is not found.
// Returns a *NotFoundError when no entities are found.
func (wfq *WithoutFieldsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{withoutfields.Label}
	default:
		err = &NotSingularError{withoutfields.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) OnlyIDX(ctx context.Context) int {
	id, err := wfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WithoutFieldsSlice.
func (wfq *WithoutFieldsQuery) All(ctx context.Context) ([]*WithoutFields, error) {
	if err := wfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) AllX(ctx context.Context) []*WithoutFields {
	nodes, err := wfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WithoutFields IDs.
func (wfq *WithoutFieldsQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wfq.Select(withoutfields.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) IDsX(ctx context.Context) []int {
	ids, err := wfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wfq *WithoutFieldsQuery) Count(ctx context.Context) (int, error) {
	if err := wfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) CountX(ctx context.Context) int {
	count, err := wfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wfq *WithoutFieldsQuery) Exist(ctx context.Context) (bool, error) {
	if err := wfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wfq *WithoutFieldsQuery) ExistX(ctx context.Context) bool {
	exist, err := wfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WithoutFieldsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wfq *WithoutFieldsQuery) Clone() *WithoutFieldsQuery {
	if wfq == nil {
		return nil
	}
	return &WithoutFieldsQuery{
		config:     wfq.config,
		limit:      wfq.limit,
		offset:     wfq.offset,
		order:      append([]OrderFunc{}, wfq.order...),
		predicates: append([]predicate.WithoutFields{}, wfq.predicates...),
		// clone intermediate query.
		sql:  wfq.sql.Clone(),
		path: wfq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (wfq *WithoutFieldsQuery) GroupBy(field string, fields ...string) *WithoutFieldsGroupBy {
	group := &WithoutFieldsGroupBy{config: wfq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wfq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (wfq *WithoutFieldsQuery) Select(field string, fields ...string) *WithoutFieldsSelect {
	wfq.fields = append([]string{field}, fields...)
	return &WithoutFieldsSelect{WithoutFieldsQuery: wfq}
}

func (wfq *WithoutFieldsQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wfq.fields {
		if !withoutfields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wfq.path != nil {
		prev, err := wfq.path(ctx)
		if err != nil {
			return err
		}
		wfq.sql = prev
	}
	return nil
}

func (wfq *WithoutFieldsQuery) sqlAll(ctx context.Context) ([]*WithoutFields, error) {
	var (
		nodes = []*WithoutFields{}
		_spec = wfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &WithoutFields{config: wfq.config}
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
	if err := sqlgraph.QueryNodes(ctx, wfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wfq *WithoutFieldsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wfq.querySpec()
	return sqlgraph.CountNodes(ctx, wfq.driver, _spec)
}

func (wfq *WithoutFieldsQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wfq *WithoutFieldsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withoutfields.Table,
			Columns: withoutfields.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: withoutfields.FieldID,
			},
		},
		From:   wfq.sql,
		Unique: true,
	}
	if unique := wfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withoutfields.FieldID)
		for i := range fields {
			if fields[i] != withoutfields.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wfq *WithoutFieldsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wfq.driver.Dialect())
	t1 := builder.Table(withoutfields.Table)
	selector := builder.Select(t1.Columns(withoutfields.Columns...)...).From(t1)
	if wfq.sql != nil {
		selector = wfq.sql
		selector.Select(selector.Columns(withoutfields.Columns...)...)
	}
	for _, p := range wfq.predicates {
		p(selector)
	}
	for _, p := range wfq.order {
		p(selector)
	}
	if offset := wfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithoutFieldsGroupBy is the group-by builder for WithoutFields entities.
type WithoutFieldsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wfgb *WithoutFieldsGroupBy) Aggregate(fns ...AggregateFunc) *WithoutFieldsGroupBy {
	wfgb.fns = append(wfgb.fns, fns...)
	return wfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wfgb *WithoutFieldsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wfgb.path(ctx)
	if err != nil {
		return err
	}
	wfgb.sql = query
	return wfgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wfgb *WithoutFieldsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wfgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfgb *WithoutFieldsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wfgb.fields) > 1 {
		return nil, errors.New("ent: WithoutFieldsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wfgb *WithoutFieldsGroupBy) StringsX(ctx context.Context) []string {
	v, err := wfgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfgb *WithoutFieldsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wfgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withoutfields.Label}
	default:
		err = fmt.Errorf("ent: WithoutFieldsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wfgb *WithoutFieldsGroupBy) StringX(ctx context.Context) string {
	v, err := wfgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfgb *WithoutFieldsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wfgb.fields) > 1 {
		return nil, errors.New("ent: WithoutFieldsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wfgb *WithoutFieldsGroupBy) IntsX(ctx context.Context) []int {
	v, err := wfgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfgb *WithoutFieldsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wfgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withoutfields.Label}
	default:
		err = fmt.Errorf("ent: WithoutFieldsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wfgb *WithoutFieldsGroupBy) IntX(ctx context.Context) int {
	v, err := wfgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfgb *WithoutFieldsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wfgb.fields) > 1 {
		return nil, errors.New("ent: WithoutFieldsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wfgb *WithoutFieldsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wfgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfgb *WithoutFieldsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wfgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withoutfields.Label}
	default:
		err = fmt.Errorf("ent: WithoutFieldsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wfgb *WithoutFieldsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wfgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wfgb *WithoutFieldsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wfgb.fields) > 1 {
		return nil, errors.New("ent: WithoutFieldsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wfgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wfgb *WithoutFieldsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wfgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wfgb *WithoutFieldsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wfgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withoutfields.Label}
	default:
		err = fmt.Errorf("ent: WithoutFieldsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wfgb *WithoutFieldsGroupBy) BoolX(ctx context.Context) bool {
	v, err := wfgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wfgb *WithoutFieldsGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wfgb.fields {
		if !withoutfields.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wfgb *WithoutFieldsGroupBy) sqlQuery() *sql.Selector {
	selector := wfgb.sql
	columns := make([]string, 0, len(wfgb.fields)+len(wfgb.fns))
	columns = append(columns, wfgb.fields...)
	for _, fn := range wfgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(wfgb.fields...)
}

// WithoutFieldsSelect is the builder for selecting fields of WithoutFields entities.
type WithoutFieldsSelect struct {
	*WithoutFieldsQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wfs *WithoutFieldsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wfs.prepareQuery(ctx); err != nil {
		return err
	}
	wfs.sql = wfs.WithoutFieldsQuery.sqlQuery(ctx)
	return wfs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wfs *WithoutFieldsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wfs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (wfs *WithoutFieldsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wfs.fields) > 1 {
		return nil, errors.New("ent: WithoutFieldsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wfs *WithoutFieldsSelect) StringsX(ctx context.Context) []string {
	v, err := wfs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (wfs *WithoutFieldsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wfs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withoutfields.Label}
	default:
		err = fmt.Errorf("ent: WithoutFieldsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wfs *WithoutFieldsSelect) StringX(ctx context.Context) string {
	v, err := wfs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (wfs *WithoutFieldsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wfs.fields) > 1 {
		return nil, errors.New("ent: WithoutFieldsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wfs *WithoutFieldsSelect) IntsX(ctx context.Context) []int {
	v, err := wfs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (wfs *WithoutFieldsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wfs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withoutfields.Label}
	default:
		err = fmt.Errorf("ent: WithoutFieldsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wfs *WithoutFieldsSelect) IntX(ctx context.Context) int {
	v, err := wfs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (wfs *WithoutFieldsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wfs.fields) > 1 {
		return nil, errors.New("ent: WithoutFieldsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wfs *WithoutFieldsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wfs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (wfs *WithoutFieldsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wfs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withoutfields.Label}
	default:
		err = fmt.Errorf("ent: WithoutFieldsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wfs *WithoutFieldsSelect) Float64X(ctx context.Context) float64 {
	v, err := wfs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (wfs *WithoutFieldsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wfs.fields) > 1 {
		return nil, errors.New("ent: WithoutFieldsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wfs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wfs *WithoutFieldsSelect) BoolsX(ctx context.Context) []bool {
	v, err := wfs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (wfs *WithoutFieldsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wfs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{withoutfields.Label}
	default:
		err = fmt.Errorf("ent: WithoutFieldsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wfs *WithoutFieldsSelect) BoolX(ctx context.Context) bool {
	v, err := wfs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wfs *WithoutFieldsSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wfs.sqlQuery().Query()
	if err := wfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wfs *WithoutFieldsSelect) sqlQuery() sql.Querier {
	selector := wfs.sql
	selector.Select(selector.Columns(wfs.fields...)...)
	return selector
}
