// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Jamescog/birth-buddy/ent/predicate"
	"github.com/Jamescog/birth-buddy/ent/users"
)

// UsersUpdate is the builder for updating Users entities.
type UsersUpdate struct {
	config
	hooks    []Hook
	mutation *UsersMutation
}

// Where appends a list predicates to the UsersUpdate builder.
func (uu *UsersUpdate) Where(ps ...predicate.Users) *UsersUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetTelegramID sets the "telegram_id" field.
func (uu *UsersUpdate) SetTelegramID(i int) *UsersUpdate {
	uu.mutation.ResetTelegramID()
	uu.mutation.SetTelegramID(i)
	return uu
}

// SetNillableTelegramID sets the "telegram_id" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableTelegramID(i *int) *UsersUpdate {
	if i != nil {
		uu.SetTelegramID(*i)
	}
	return uu
}

// AddTelegramID adds i to the "telegram_id" field.
func (uu *UsersUpdate) AddTelegramID(i int) *UsersUpdate {
	uu.mutation.AddTelegramID(i)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UsersUpdate) SetUsername(s string) *UsersUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableUsername(s *string) *UsersUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// ClearUsername clears the value of the "username" field.
func (uu *UsersUpdate) ClearUsername() *UsersUpdate {
	uu.mutation.ClearUsername()
	return uu
}

// SetFullName sets the "full_name" field.
func (uu *UsersUpdate) SetFullName(s string) *UsersUpdate {
	uu.mutation.SetFullName(s)
	return uu
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableFullName(s *string) *UsersUpdate {
	if s != nil {
		uu.SetFullName(*s)
	}
	return uu
}

// SetIsPremium sets the "is_premium" field.
func (uu *UsersUpdate) SetIsPremium(b bool) *UsersUpdate {
	uu.mutation.SetIsPremium(b)
	return uu
}

// SetNillableIsPremium sets the "is_premium" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableIsPremium(b *bool) *UsersUpdate {
	if b != nil {
		uu.SetIsPremium(*b)
	}
	return uu
}

// SetBirthdayCount sets the "birthday_count" field.
func (uu *UsersUpdate) SetBirthdayCount(i int) *UsersUpdate {
	uu.mutation.ResetBirthdayCount()
	uu.mutation.SetBirthdayCount(i)
	return uu
}

// SetNillableBirthdayCount sets the "birthday_count" field if the given value is not nil.
func (uu *UsersUpdate) SetNillableBirthdayCount(i *int) *UsersUpdate {
	if i != nil {
		uu.SetBirthdayCount(*i)
	}
	return uu
}

// AddBirthdayCount adds i to the "birthday_count" field.
func (uu *UsersUpdate) AddBirthdayCount(i int) *UsersUpdate {
	uu.mutation.AddBirthdayCount(i)
	return uu
}

// Mutation returns the UsersMutation object of the builder.
func (uu *UsersUpdate) Mutation() *UsersMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UsersUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UsersUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UsersUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UsersUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UsersUpdate) check() error {
	if v, ok := uu.mutation.TelegramID(); ok {
		if err := users.TelegramIDValidator(v); err != nil {
			return &ValidationError{Name: "telegram_id", err: fmt.Errorf(`ent: validator failed for field "Users.telegram_id": %w`, err)}
		}
	}
	if v, ok := uu.mutation.BirthdayCount(); ok {
		if err := users.BirthdayCountValidator(v); err != nil {
			return &ValidationError{Name: "birthday_count", err: fmt.Errorf(`ent: validator failed for field "Users.birthday_count": %w`, err)}
		}
	}
	return nil
}

func (uu *UsersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.TelegramID(); ok {
		_spec.SetField(users.FieldTelegramID, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedTelegramID(); ok {
		_spec.AddField(users.FieldTelegramID, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
	}
	if uu.mutation.UsernameCleared() {
		_spec.ClearField(users.FieldUsername, field.TypeString)
	}
	if value, ok := uu.mutation.FullName(); ok {
		_spec.SetField(users.FieldFullName, field.TypeString, value)
	}
	if value, ok := uu.mutation.IsPremium(); ok {
		_spec.SetField(users.FieldIsPremium, field.TypeBool, value)
	}
	if value, ok := uu.mutation.BirthdayCount(); ok {
		_spec.SetField(users.FieldBirthdayCount, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedBirthdayCount(); ok {
		_spec.AddField(users.FieldBirthdayCount, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UsersUpdateOne is the builder for updating a single Users entity.
type UsersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsersMutation
}

// SetTelegramID sets the "telegram_id" field.
func (uuo *UsersUpdateOne) SetTelegramID(i int) *UsersUpdateOne {
	uuo.mutation.ResetTelegramID()
	uuo.mutation.SetTelegramID(i)
	return uuo
}

// SetNillableTelegramID sets the "telegram_id" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableTelegramID(i *int) *UsersUpdateOne {
	if i != nil {
		uuo.SetTelegramID(*i)
	}
	return uuo
}

// AddTelegramID adds i to the "telegram_id" field.
func (uuo *UsersUpdateOne) AddTelegramID(i int) *UsersUpdateOne {
	uuo.mutation.AddTelegramID(i)
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UsersUpdateOne) SetUsername(s string) *UsersUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableUsername(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// ClearUsername clears the value of the "username" field.
func (uuo *UsersUpdateOne) ClearUsername() *UsersUpdateOne {
	uuo.mutation.ClearUsername()
	return uuo
}

// SetFullName sets the "full_name" field.
func (uuo *UsersUpdateOne) SetFullName(s string) *UsersUpdateOne {
	uuo.mutation.SetFullName(s)
	return uuo
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableFullName(s *string) *UsersUpdateOne {
	if s != nil {
		uuo.SetFullName(*s)
	}
	return uuo
}

// SetIsPremium sets the "is_premium" field.
func (uuo *UsersUpdateOne) SetIsPremium(b bool) *UsersUpdateOne {
	uuo.mutation.SetIsPremium(b)
	return uuo
}

// SetNillableIsPremium sets the "is_premium" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableIsPremium(b *bool) *UsersUpdateOne {
	if b != nil {
		uuo.SetIsPremium(*b)
	}
	return uuo
}

// SetBirthdayCount sets the "birthday_count" field.
func (uuo *UsersUpdateOne) SetBirthdayCount(i int) *UsersUpdateOne {
	uuo.mutation.ResetBirthdayCount()
	uuo.mutation.SetBirthdayCount(i)
	return uuo
}

// SetNillableBirthdayCount sets the "birthday_count" field if the given value is not nil.
func (uuo *UsersUpdateOne) SetNillableBirthdayCount(i *int) *UsersUpdateOne {
	if i != nil {
		uuo.SetBirthdayCount(*i)
	}
	return uuo
}

// AddBirthdayCount adds i to the "birthday_count" field.
func (uuo *UsersUpdateOne) AddBirthdayCount(i int) *UsersUpdateOne {
	uuo.mutation.AddBirthdayCount(i)
	return uuo
}

// Mutation returns the UsersMutation object of the builder.
func (uuo *UsersUpdateOne) Mutation() *UsersMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UsersUpdate builder.
func (uuo *UsersUpdateOne) Where(ps ...predicate.Users) *UsersUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UsersUpdateOne) Select(field string, fields ...string) *UsersUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Users entity.
func (uuo *UsersUpdateOne) Save(ctx context.Context) (*Users, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UsersUpdateOne) SaveX(ctx context.Context) *Users {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UsersUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UsersUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UsersUpdateOne) check() error {
	if v, ok := uuo.mutation.TelegramID(); ok {
		if err := users.TelegramIDValidator(v); err != nil {
			return &ValidationError{Name: "telegram_id", err: fmt.Errorf(`ent: validator failed for field "Users.telegram_id": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.BirthdayCount(); ok {
		if err := users.BirthdayCountValidator(v); err != nil {
			return &ValidationError{Name: "birthday_count", err: fmt.Errorf(`ent: validator failed for field "Users.birthday_count": %w`, err)}
		}
	}
	return nil
}

func (uuo *UsersUpdateOne) sqlSave(ctx context.Context) (_node *Users, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Users.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, users.FieldID)
		for _, f := range fields {
			if !users.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != users.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.TelegramID(); ok {
		_spec.SetField(users.FieldTelegramID, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedTelegramID(); ok {
		_spec.AddField(users.FieldTelegramID, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
	}
	if uuo.mutation.UsernameCleared() {
		_spec.ClearField(users.FieldUsername, field.TypeString)
	}
	if value, ok := uuo.mutation.FullName(); ok {
		_spec.SetField(users.FieldFullName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.IsPremium(); ok {
		_spec.SetField(users.FieldIsPremium, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.BirthdayCount(); ok {
		_spec.SetField(users.FieldBirthdayCount, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedBirthdayCount(); ok {
		_spec.AddField(users.FieldBirthdayCount, field.TypeInt, value)
	}
	_node = &Users{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}