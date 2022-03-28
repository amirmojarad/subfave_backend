// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"interface_project/ent/searchkeyword"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SearchKeywordCreate is the builder for creating a SearchKeyword entity.
type SearchKeywordCreate struct {
	config
	mutation *SearchKeywordMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (skc *SearchKeywordCreate) SetTitle(s string) *SearchKeywordCreate {
	skc.mutation.SetTitle(s)
	return skc
}

// SetRate sets the "rate" field.
func (skc *SearchKeywordCreate) SetRate(u uint16) *SearchKeywordCreate {
	skc.mutation.SetRate(u)
	return skc
}

// SetNillableRate sets the "rate" field if the given value is not nil.
func (skc *SearchKeywordCreate) SetNillableRate(u *uint16) *SearchKeywordCreate {
	if u != nil {
		skc.SetRate(*u)
	}
	return skc
}

// Mutation returns the SearchKeywordMutation object of the builder.
func (skc *SearchKeywordCreate) Mutation() *SearchKeywordMutation {
	return skc.mutation
}

// Save creates the SearchKeyword in the database.
func (skc *SearchKeywordCreate) Save(ctx context.Context) (*SearchKeyword, error) {
	var (
		err  error
		node *SearchKeyword
	)
	skc.defaults()
	if len(skc.hooks) == 0 {
		if err = skc.check(); err != nil {
			return nil, err
		}
		node, err = skc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SearchKeywordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = skc.check(); err != nil {
				return nil, err
			}
			skc.mutation = mutation
			if node, err = skc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(skc.hooks) - 1; i >= 0; i-- {
			if skc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = skc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, skc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (skc *SearchKeywordCreate) SaveX(ctx context.Context) *SearchKeyword {
	v, err := skc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (skc *SearchKeywordCreate) Exec(ctx context.Context) error {
	_, err := skc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (skc *SearchKeywordCreate) ExecX(ctx context.Context) {
	if err := skc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (skc *SearchKeywordCreate) defaults() {
	if _, ok := skc.mutation.Rate(); !ok {
		v := searchkeyword.DefaultRate
		skc.mutation.SetRate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (skc *SearchKeywordCreate) check() error {
	if _, ok := skc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "SearchKeyword.title"`)}
	}
	if v, ok := skc.mutation.Title(); ok {
		if err := searchkeyword.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "SearchKeyword.title": %w`, err)}
		}
	}
	if _, ok := skc.mutation.Rate(); !ok {
		return &ValidationError{Name: "rate", err: errors.New(`ent: missing required field "SearchKeyword.rate"`)}
	}
	return nil
}

func (skc *SearchKeywordCreate) sqlSave(ctx context.Context) (*SearchKeyword, error) {
	_node, _spec := skc.createSpec()
	if err := sqlgraph.CreateNode(ctx, skc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (skc *SearchKeywordCreate) createSpec() (*SearchKeyword, *sqlgraph.CreateSpec) {
	var (
		_node = &SearchKeyword{config: skc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: searchkeyword.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: searchkeyword.FieldID,
			},
		}
	)
	if value, ok := skc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: searchkeyword.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := skc.mutation.Rate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint16,
			Value:  value,
			Column: searchkeyword.FieldRate,
		})
		_node.Rate = value
	}
	return _node, _spec
}

// SearchKeywordCreateBulk is the builder for creating many SearchKeyword entities in bulk.
type SearchKeywordCreateBulk struct {
	config
	builders []*SearchKeywordCreate
}

// Save creates the SearchKeyword entities in the database.
func (skcb *SearchKeywordCreateBulk) Save(ctx context.Context) ([]*SearchKeyword, error) {
	specs := make([]*sqlgraph.CreateSpec, len(skcb.builders))
	nodes := make([]*SearchKeyword, len(skcb.builders))
	mutators := make([]Mutator, len(skcb.builders))
	for i := range skcb.builders {
		func(i int, root context.Context) {
			builder := skcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SearchKeywordMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, skcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, skcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, skcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (skcb *SearchKeywordCreateBulk) SaveX(ctx context.Context) []*SearchKeyword {
	v, err := skcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (skcb *SearchKeywordCreateBulk) Exec(ctx context.Context) error {
	_, err := skcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (skcb *SearchKeywordCreateBulk) ExecX(ctx context.Context) {
	if err := skcb.Exec(ctx); err != nil {
		panic(err)
	}
}