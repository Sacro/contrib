// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithoptionals"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MessageWithOptionalsUpdate is the builder for updating MessageWithOptionals entities.
type MessageWithOptionalsUpdate struct {
	config
	hooks    []Hook
	mutation *MessageWithOptionalsMutation
}

// Where adds a new predicate for the MessageWithOptionalsUpdate builder.
func (mwou *MessageWithOptionalsUpdate) Where(ps ...predicate.MessageWithOptionals) *MessageWithOptionalsUpdate {
	mwou.mutation.predicates = append(mwou.mutation.predicates, ps...)
	return mwou
}

// SetStrOptional sets the "str_optional" field.
func (mwou *MessageWithOptionalsUpdate) SetStrOptional(s string) *MessageWithOptionalsUpdate {
	mwou.mutation.SetStrOptional(s)
	return mwou
}

// SetNillableStrOptional sets the "str_optional" field if the given value is not nil.
func (mwou *MessageWithOptionalsUpdate) SetNillableStrOptional(s *string) *MessageWithOptionalsUpdate {
	if s != nil {
		mwou.SetStrOptional(*s)
	}
	return mwou
}

// ClearStrOptional clears the value of the "str_optional" field.
func (mwou *MessageWithOptionalsUpdate) ClearStrOptional() *MessageWithOptionalsUpdate {
	mwou.mutation.ClearStrOptional()
	return mwou
}

// SetIntOptional sets the "int_optional" field.
func (mwou *MessageWithOptionalsUpdate) SetIntOptional(i int8) *MessageWithOptionalsUpdate {
	mwou.mutation.ResetIntOptional()
	mwou.mutation.SetIntOptional(i)
	return mwou
}

// SetNillableIntOptional sets the "int_optional" field if the given value is not nil.
func (mwou *MessageWithOptionalsUpdate) SetNillableIntOptional(i *int8) *MessageWithOptionalsUpdate {
	if i != nil {
		mwou.SetIntOptional(*i)
	}
	return mwou
}

// AddIntOptional adds i to the "int_optional" field.
func (mwou *MessageWithOptionalsUpdate) AddIntOptional(i int8) *MessageWithOptionalsUpdate {
	mwou.mutation.AddIntOptional(i)
	return mwou
}

// ClearIntOptional clears the value of the "int_optional" field.
func (mwou *MessageWithOptionalsUpdate) ClearIntOptional() *MessageWithOptionalsUpdate {
	mwou.mutation.ClearIntOptional()
	return mwou
}

// SetUintOptional sets the "uint_optional" field.
func (mwou *MessageWithOptionalsUpdate) SetUintOptional(u uint8) *MessageWithOptionalsUpdate {
	mwou.mutation.ResetUintOptional()
	mwou.mutation.SetUintOptional(u)
	return mwou
}

// SetNillableUintOptional sets the "uint_optional" field if the given value is not nil.
func (mwou *MessageWithOptionalsUpdate) SetNillableUintOptional(u *uint8) *MessageWithOptionalsUpdate {
	if u != nil {
		mwou.SetUintOptional(*u)
	}
	return mwou
}

// AddUintOptional adds u to the "uint_optional" field.
func (mwou *MessageWithOptionalsUpdate) AddUintOptional(u uint8) *MessageWithOptionalsUpdate {
	mwou.mutation.AddUintOptional(u)
	return mwou
}

// ClearUintOptional clears the value of the "uint_optional" field.
func (mwou *MessageWithOptionalsUpdate) ClearUintOptional() *MessageWithOptionalsUpdate {
	mwou.mutation.ClearUintOptional()
	return mwou
}

// SetFloatOptional sets the "float_optional" field.
func (mwou *MessageWithOptionalsUpdate) SetFloatOptional(f float32) *MessageWithOptionalsUpdate {
	mwou.mutation.ResetFloatOptional()
	mwou.mutation.SetFloatOptional(f)
	return mwou
}

// SetNillableFloatOptional sets the "float_optional" field if the given value is not nil.
func (mwou *MessageWithOptionalsUpdate) SetNillableFloatOptional(f *float32) *MessageWithOptionalsUpdate {
	if f != nil {
		mwou.SetFloatOptional(*f)
	}
	return mwou
}

// AddFloatOptional adds f to the "float_optional" field.
func (mwou *MessageWithOptionalsUpdate) AddFloatOptional(f float32) *MessageWithOptionalsUpdate {
	mwou.mutation.AddFloatOptional(f)
	return mwou
}

// ClearFloatOptional clears the value of the "float_optional" field.
func (mwou *MessageWithOptionalsUpdate) ClearFloatOptional() *MessageWithOptionalsUpdate {
	mwou.mutation.ClearFloatOptional()
	return mwou
}

// SetBoolOptional sets the "bool_optional" field.
func (mwou *MessageWithOptionalsUpdate) SetBoolOptional(b bool) *MessageWithOptionalsUpdate {
	mwou.mutation.SetBoolOptional(b)
	return mwou
}

// SetNillableBoolOptional sets the "bool_optional" field if the given value is not nil.
func (mwou *MessageWithOptionalsUpdate) SetNillableBoolOptional(b *bool) *MessageWithOptionalsUpdate {
	if b != nil {
		mwou.SetBoolOptional(*b)
	}
	return mwou
}

// ClearBoolOptional clears the value of the "bool_optional" field.
func (mwou *MessageWithOptionalsUpdate) ClearBoolOptional() *MessageWithOptionalsUpdate {
	mwou.mutation.ClearBoolOptional()
	return mwou
}

// SetBytesOptional sets the "bytes_optional" field.
func (mwou *MessageWithOptionalsUpdate) SetBytesOptional(b []byte) *MessageWithOptionalsUpdate {
	mwou.mutation.SetBytesOptional(b)
	return mwou
}

// ClearBytesOptional clears the value of the "bytes_optional" field.
func (mwou *MessageWithOptionalsUpdate) ClearBytesOptional() *MessageWithOptionalsUpdate {
	mwou.mutation.ClearBytesOptional()
	return mwou
}

// SetUUIDOptional sets the "uuid_optional" field.
func (mwou *MessageWithOptionalsUpdate) SetUUIDOptional(u uuid.UUID) *MessageWithOptionalsUpdate {
	mwou.mutation.SetUUIDOptional(u)
	return mwou
}

// ClearUUIDOptional clears the value of the "uuid_optional" field.
func (mwou *MessageWithOptionalsUpdate) ClearUUIDOptional() *MessageWithOptionalsUpdate {
	mwou.mutation.ClearUUIDOptional()
	return mwou
}

// SetTimeOptional sets the "time_optional" field.
func (mwou *MessageWithOptionalsUpdate) SetTimeOptional(t time.Time) *MessageWithOptionalsUpdate {
	mwou.mutation.SetTimeOptional(t)
	return mwou
}

// SetNillableTimeOptional sets the "time_optional" field if the given value is not nil.
func (mwou *MessageWithOptionalsUpdate) SetNillableTimeOptional(t *time.Time) *MessageWithOptionalsUpdate {
	if t != nil {
		mwou.SetTimeOptional(*t)
	}
	return mwou
}

// ClearTimeOptional clears the value of the "time_optional" field.
func (mwou *MessageWithOptionalsUpdate) ClearTimeOptional() *MessageWithOptionalsUpdate {
	mwou.mutation.ClearTimeOptional()
	return mwou
}

// Mutation returns the MessageWithOptionalsMutation object of the builder.
func (mwou *MessageWithOptionalsUpdate) Mutation() *MessageWithOptionalsMutation {
	return mwou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mwou *MessageWithOptionalsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mwou.hooks) == 0 {
		affected, err = mwou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithOptionalsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mwou.mutation = mutation
			affected, err = mwou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mwou.hooks) - 1; i >= 0; i-- {
			mut = mwou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mwou *MessageWithOptionalsUpdate) SaveX(ctx context.Context) int {
	affected, err := mwou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mwou *MessageWithOptionalsUpdate) Exec(ctx context.Context) error {
	_, err := mwou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwou *MessageWithOptionalsUpdate) ExecX(ctx context.Context) {
	if err := mwou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwou *MessageWithOptionalsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagewithoptionals.Table,
			Columns: messagewithoptionals.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithoptionals.FieldID,
			},
		},
	}
	if ps := mwou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mwou.mutation.StrOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messagewithoptionals.FieldStrOptional,
		})
	}
	if mwou.mutation.StrOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: messagewithoptionals.FieldStrOptional,
		})
	}
	if value, ok := mwou.mutation.IntOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: messagewithoptionals.FieldIntOptional,
		})
	}
	if value, ok := mwou.mutation.AddedIntOptional(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: messagewithoptionals.FieldIntOptional,
		})
	}
	if mwou.mutation.IntOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Column: messagewithoptionals.FieldIntOptional,
		})
	}
	if value, ok := mwou.mutation.UintOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: messagewithoptionals.FieldUintOptional,
		})
	}
	if value, ok := mwou.mutation.AddedUintOptional(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: messagewithoptionals.FieldUintOptional,
		})
	}
	if mwou.mutation.UintOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: messagewithoptionals.FieldUintOptional,
		})
	}
	if value, ok := mwou.mutation.FloatOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: messagewithoptionals.FieldFloatOptional,
		})
	}
	if value, ok := mwou.mutation.AddedFloatOptional(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: messagewithoptionals.FieldFloatOptional,
		})
	}
	if mwou.mutation.FloatOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Column: messagewithoptionals.FieldFloatOptional,
		})
	}
	if value, ok := mwou.mutation.BoolOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: messagewithoptionals.FieldBoolOptional,
		})
	}
	if mwou.mutation.BoolOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: messagewithoptionals.FieldBoolOptional,
		})
	}
	if value, ok := mwou.mutation.BytesOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: messagewithoptionals.FieldBytesOptional,
		})
	}
	if mwou.mutation.BytesOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: messagewithoptionals.FieldBytesOptional,
		})
	}
	if value, ok := mwou.mutation.UUIDOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: messagewithoptionals.FieldUUIDOptional,
		})
	}
	if mwou.mutation.UUIDOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: messagewithoptionals.FieldUUIDOptional,
		})
	}
	if value, ok := mwou.mutation.TimeOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messagewithoptionals.FieldTimeOptional,
		})
	}
	if mwou.mutation.TimeOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: messagewithoptionals.FieldTimeOptional,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mwou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithoptionals.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MessageWithOptionalsUpdateOne is the builder for updating a single MessageWithOptionals entity.
type MessageWithOptionalsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageWithOptionalsMutation
}

// SetStrOptional sets the "str_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) SetStrOptional(s string) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.SetStrOptional(s)
	return mwouo
}

// SetNillableStrOptional sets the "str_optional" field if the given value is not nil.
func (mwouo *MessageWithOptionalsUpdateOne) SetNillableStrOptional(s *string) *MessageWithOptionalsUpdateOne {
	if s != nil {
		mwouo.SetStrOptional(*s)
	}
	return mwouo
}

// ClearStrOptional clears the value of the "str_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) ClearStrOptional() *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ClearStrOptional()
	return mwouo
}

// SetIntOptional sets the "int_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) SetIntOptional(i int8) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ResetIntOptional()
	mwouo.mutation.SetIntOptional(i)
	return mwouo
}

// SetNillableIntOptional sets the "int_optional" field if the given value is not nil.
func (mwouo *MessageWithOptionalsUpdateOne) SetNillableIntOptional(i *int8) *MessageWithOptionalsUpdateOne {
	if i != nil {
		mwouo.SetIntOptional(*i)
	}
	return mwouo
}

// AddIntOptional adds i to the "int_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) AddIntOptional(i int8) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.AddIntOptional(i)
	return mwouo
}

// ClearIntOptional clears the value of the "int_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) ClearIntOptional() *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ClearIntOptional()
	return mwouo
}

// SetUintOptional sets the "uint_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) SetUintOptional(u uint8) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ResetUintOptional()
	mwouo.mutation.SetUintOptional(u)
	return mwouo
}

// SetNillableUintOptional sets the "uint_optional" field if the given value is not nil.
func (mwouo *MessageWithOptionalsUpdateOne) SetNillableUintOptional(u *uint8) *MessageWithOptionalsUpdateOne {
	if u != nil {
		mwouo.SetUintOptional(*u)
	}
	return mwouo
}

// AddUintOptional adds u to the "uint_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) AddUintOptional(u uint8) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.AddUintOptional(u)
	return mwouo
}

// ClearUintOptional clears the value of the "uint_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) ClearUintOptional() *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ClearUintOptional()
	return mwouo
}

// SetFloatOptional sets the "float_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) SetFloatOptional(f float32) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ResetFloatOptional()
	mwouo.mutation.SetFloatOptional(f)
	return mwouo
}

// SetNillableFloatOptional sets the "float_optional" field if the given value is not nil.
func (mwouo *MessageWithOptionalsUpdateOne) SetNillableFloatOptional(f *float32) *MessageWithOptionalsUpdateOne {
	if f != nil {
		mwouo.SetFloatOptional(*f)
	}
	return mwouo
}

// AddFloatOptional adds f to the "float_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) AddFloatOptional(f float32) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.AddFloatOptional(f)
	return mwouo
}

// ClearFloatOptional clears the value of the "float_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) ClearFloatOptional() *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ClearFloatOptional()
	return mwouo
}

// SetBoolOptional sets the "bool_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) SetBoolOptional(b bool) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.SetBoolOptional(b)
	return mwouo
}

// SetNillableBoolOptional sets the "bool_optional" field if the given value is not nil.
func (mwouo *MessageWithOptionalsUpdateOne) SetNillableBoolOptional(b *bool) *MessageWithOptionalsUpdateOne {
	if b != nil {
		mwouo.SetBoolOptional(*b)
	}
	return mwouo
}

// ClearBoolOptional clears the value of the "bool_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) ClearBoolOptional() *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ClearBoolOptional()
	return mwouo
}

// SetBytesOptional sets the "bytes_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) SetBytesOptional(b []byte) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.SetBytesOptional(b)
	return mwouo
}

// ClearBytesOptional clears the value of the "bytes_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) ClearBytesOptional() *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ClearBytesOptional()
	return mwouo
}

// SetUUIDOptional sets the "uuid_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) SetUUIDOptional(u uuid.UUID) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.SetUUIDOptional(u)
	return mwouo
}

// ClearUUIDOptional clears the value of the "uuid_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) ClearUUIDOptional() *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ClearUUIDOptional()
	return mwouo
}

// SetTimeOptional sets the "time_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) SetTimeOptional(t time.Time) *MessageWithOptionalsUpdateOne {
	mwouo.mutation.SetTimeOptional(t)
	return mwouo
}

// SetNillableTimeOptional sets the "time_optional" field if the given value is not nil.
func (mwouo *MessageWithOptionalsUpdateOne) SetNillableTimeOptional(t *time.Time) *MessageWithOptionalsUpdateOne {
	if t != nil {
		mwouo.SetTimeOptional(*t)
	}
	return mwouo
}

// ClearTimeOptional clears the value of the "time_optional" field.
func (mwouo *MessageWithOptionalsUpdateOne) ClearTimeOptional() *MessageWithOptionalsUpdateOne {
	mwouo.mutation.ClearTimeOptional()
	return mwouo
}

// Mutation returns the MessageWithOptionalsMutation object of the builder.
func (mwouo *MessageWithOptionalsUpdateOne) Mutation() *MessageWithOptionalsMutation {
	return mwouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mwouo *MessageWithOptionalsUpdateOne) Select(field string, fields ...string) *MessageWithOptionalsUpdateOne {
	mwouo.fields = append([]string{field}, fields...)
	return mwouo
}

// Save executes the query and returns the updated MessageWithOptionals entity.
func (mwouo *MessageWithOptionalsUpdateOne) Save(ctx context.Context) (*MessageWithOptionals, error) {
	var (
		err  error
		node *MessageWithOptionals
	)
	if len(mwouo.hooks) == 0 {
		node, err = mwouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithOptionalsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mwouo.mutation = mutation
			node, err = mwouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mwouo.hooks) - 1; i >= 0; i-- {
			mut = mwouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mwouo *MessageWithOptionalsUpdateOne) SaveX(ctx context.Context) *MessageWithOptionals {
	node, err := mwouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mwouo *MessageWithOptionalsUpdateOne) Exec(ctx context.Context) error {
	_, err := mwouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwouo *MessageWithOptionalsUpdateOne) ExecX(ctx context.Context) {
	if err := mwouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwouo *MessageWithOptionalsUpdateOne) sqlSave(ctx context.Context) (_node *MessageWithOptionals, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagewithoptionals.Table,
			Columns: messagewithoptionals.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithoptionals.FieldID,
			},
		},
	}
	id, ok := mwouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing MessageWithOptionals.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := mwouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithoptionals.FieldID)
		for _, f := range fields {
			if !messagewithoptionals.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messagewithoptionals.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mwouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mwouo.mutation.StrOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: messagewithoptionals.FieldStrOptional,
		})
	}
	if mwouo.mutation.StrOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: messagewithoptionals.FieldStrOptional,
		})
	}
	if value, ok := mwouo.mutation.IntOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: messagewithoptionals.FieldIntOptional,
		})
	}
	if value, ok := mwouo.mutation.AddedIntOptional(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: messagewithoptionals.FieldIntOptional,
		})
	}
	if mwouo.mutation.IntOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Column: messagewithoptionals.FieldIntOptional,
		})
	}
	if value, ok := mwouo.mutation.UintOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: messagewithoptionals.FieldUintOptional,
		})
	}
	if value, ok := mwouo.mutation.AddedUintOptional(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: messagewithoptionals.FieldUintOptional,
		})
	}
	if mwouo.mutation.UintOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: messagewithoptionals.FieldUintOptional,
		})
	}
	if value, ok := mwouo.mutation.FloatOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: messagewithoptionals.FieldFloatOptional,
		})
	}
	if value, ok := mwouo.mutation.AddedFloatOptional(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: messagewithoptionals.FieldFloatOptional,
		})
	}
	if mwouo.mutation.FloatOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Column: messagewithoptionals.FieldFloatOptional,
		})
	}
	if value, ok := mwouo.mutation.BoolOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: messagewithoptionals.FieldBoolOptional,
		})
	}
	if mwouo.mutation.BoolOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: messagewithoptionals.FieldBoolOptional,
		})
	}
	if value, ok := mwouo.mutation.BytesOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: messagewithoptionals.FieldBytesOptional,
		})
	}
	if mwouo.mutation.BytesOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: messagewithoptionals.FieldBytesOptional,
		})
	}
	if value, ok := mwouo.mutation.UUIDOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: messagewithoptionals.FieldUUIDOptional,
		})
	}
	if mwouo.mutation.UUIDOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: messagewithoptionals.FieldUUIDOptional,
		})
	}
	if value, ok := mwouo.mutation.TimeOptional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: messagewithoptionals.FieldTimeOptional,
		})
	}
	if mwouo.mutation.TimeOptionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: messagewithoptionals.FieldTimeOptional,
		})
	}
	_node = &MessageWithOptionals{config: mwouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mwouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithoptionals.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
