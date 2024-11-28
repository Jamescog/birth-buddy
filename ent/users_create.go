// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Jamescog/birth-buddy/ent/users"
)

// UsersCreate is the builder for creating a Users entity.
type UsersCreate struct {
	config
	mutation *UsersMutation
	hooks    []Hook
}

// SetTelegramID sets the "telegram_id" field.
func (uc *UsersCreate) SetTelegramID(i int) *UsersCreate {
	uc.mutation.SetTelegramID(i)
	return uc
}

// SetUsername sets the "username" field.
func (uc *UsersCreate) SetUsername(s string) *UsersCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uc *UsersCreate) SetNillableUsername(s *string) *UsersCreate {
	if s != nil {
		uc.SetUsername(*s)
	}
	return uc
}

// SetFullName sets the "full_name" field.
func (uc *UsersCreate) SetFullName(s string) *UsersCreate {
	uc.mutation.SetFullName(s)
	return uc
}

// SetIsPremium sets the "is_premium" field.
func (uc *UsersCreate) SetIsPremium(b bool) *UsersCreate {
	uc.mutation.SetIsPremium(b)
	return uc
}

// SetNillableIsPremium sets the "is_premium" field if the given value is not nil.
func (uc *UsersCreate) SetNillableIsPremium(b *bool) *UsersCreate {
	if b != nil {
		uc.SetIsPremium(*b)
	}
	return uc
}

// SetBirthdayCount sets the "birthday_count" field.
func (uc *UsersCreate) SetBirthdayCount(i int) *UsersCreate {
	uc.mutation.SetBirthdayCount(i)
	return uc
}

// Mutation returns the UsersMutation object of the builder.
func (uc *UsersCreate) Mutation() *UsersMutation {
	return uc.mutation
}

// Save creates the Users in the database.
func (uc *UsersCreate) Save(ctx context.Context) (*Users, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UsersCreate) SaveX(ctx context.Context) *Users {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UsersCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UsersCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UsersCreate) defaults() {
	if _, ok := uc.mutation.IsPremium(); !ok {
		v := users.DefaultIsPremium
		uc.mutation.SetIsPremium(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UsersCreate) check() error {
	if _, ok := uc.mutation.TelegramID(); !ok {
		return &ValidationError{Name: "telegram_id", err: errors.New(`ent: missing required field "Users.telegram_id"`)}
	}
	if v, ok := uc.mutation.TelegramID(); ok {
		if err := users.TelegramIDValidator(v); err != nil {
			return &ValidationError{Name: "telegram_id", err: fmt.Errorf(`ent: validator failed for field "Users.telegram_id": %w`, err)}
		}
	}
	if _, ok := uc.mutation.FullName(); !ok {
		return &ValidationError{Name: "full_name", err: errors.New(`ent: missing required field "Users.full_name"`)}
	}
	if _, ok := uc.mutation.IsPremium(); !ok {
		return &ValidationError{Name: "is_premium", err: errors.New(`ent: missing required field "Users.is_premium"`)}
	}
	if _, ok := uc.mutation.BirthdayCount(); !ok {
		return &ValidationError{Name: "birthday_count", err: errors.New(`ent: missing required field "Users.birthday_count"`)}
	}
	if v, ok := uc.mutation.BirthdayCount(); ok {
		if err := users.BirthdayCountValidator(v); err != nil {
			return &ValidationError{Name: "birthday_count", err: fmt.Errorf(`ent: validator failed for field "Users.birthday_count": %w`, err)}
		}
	}
	return nil
}

func (uc *UsersCreate) sqlSave(ctx context.Context) (*Users, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UsersCreate) createSpec() (*Users, *sqlgraph.CreateSpec) {
	var (
		_node = &Users{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(users.Table, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.TelegramID(); ok {
		_spec.SetField(users.FieldTelegramID, field.TypeInt, value)
		_node.TelegramID = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.FullName(); ok {
		_spec.SetField(users.FieldFullName, field.TypeString, value)
		_node.FullName = value
	}
	if value, ok := uc.mutation.IsPremium(); ok {
		_spec.SetField(users.FieldIsPremium, field.TypeBool, value)
		_node.IsPremium = value
	}
	if value, ok := uc.mutation.BirthdayCount(); ok {
		_spec.SetField(users.FieldBirthdayCount, field.TypeInt, value)
		_node.BirthdayCount = value
	}
	return _node, _spec
}

// UsersCreateBulk is the builder for creating many Users entities in bulk.
type UsersCreateBulk struct {
	config
	err      error
	builders []*UsersCreate
}

// Save creates the Users entities in the database.
func (ucb *UsersCreateBulk) Save(ctx context.Context) ([]*Users, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Users, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsersMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UsersCreateBulk) SaveX(ctx context.Context) []*Users {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UsersCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UsersCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}