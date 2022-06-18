// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"interface_project/ent/category"
	"interface_project/ent/fileentity"
	"interface_project/ent/user"
	"interface_project/ent/word"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WordCreate is the builder for creating a Word entity.
type WordCreate struct {
	config
	mutation *WordMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (wc *WordCreate) SetTitle(s string) *WordCreate {
	wc.mutation.SetTitle(s)
	return wc
}

// SetMeaning sets the "meaning" field.
func (wc *WordCreate) SetMeaning(s string) *WordCreate {
	wc.mutation.SetMeaning(s)
	return wc
}

// SetIsPreposition sets the "isPreposition" field.
func (wc *WordCreate) SetIsPreposition(b bool) *WordCreate {
	wc.mutation.SetIsPreposition(b)
	return wc
}

// SetSentence sets the "sentence" field.
func (wc *WordCreate) SetSentence(s string) *WordCreate {
	wc.mutation.SetSentence(s)
	return wc
}

// SetDuration sets the "duration" field.
func (wc *WordCreate) SetDuration(s string) *WordCreate {
	wc.mutation.SetDuration(s)
	return wc
}

// SetStart sets the "start" field.
func (wc *WordCreate) SetStart(s string) *WordCreate {
	wc.mutation.SetStart(s)
	return wc
}

// SetEnd sets the "end" field.
func (wc *WordCreate) SetEnd(s string) *WordCreate {
	wc.mutation.SetEnd(s)
	return wc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (wc *WordCreate) SetUserID(id int) *WordCreate {
	wc.mutation.SetUserID(id)
	return wc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (wc *WordCreate) SetNillableUserID(id *int) *WordCreate {
	if id != nil {
		wc = wc.SetUserID(*id)
	}
	return wc
}

// SetUser sets the "user" edge to the User entity.
func (wc *WordCreate) SetUser(u *User) *WordCreate {
	return wc.SetUserID(u.ID)
}

// SetFileID sets the "file" edge to the FileEntity entity by ID.
func (wc *WordCreate) SetFileID(id int) *WordCreate {
	wc.mutation.SetFileID(id)
	return wc
}

// SetNillableFileID sets the "file" edge to the FileEntity entity by ID if the given value is not nil.
func (wc *WordCreate) SetNillableFileID(id *int) *WordCreate {
	if id != nil {
		wc = wc.SetFileID(*id)
	}
	return wc
}

// SetFile sets the "file" edge to the FileEntity entity.
func (wc *WordCreate) SetFile(f *FileEntity) *WordCreate {
	return wc.SetFileID(f.ID)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (wc *WordCreate) AddCategoryIDs(ids ...int) *WordCreate {
	wc.mutation.AddCategoryIDs(ids...)
	return wc
}

// AddCategory adds the "category" edges to the Category entity.
func (wc *WordCreate) AddCategory(c ...*Category) *WordCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wc.AddCategoryIDs(ids...)
}

// Mutation returns the WordMutation object of the builder.
func (wc *WordCreate) Mutation() *WordMutation {
	return wc.mutation
}

// Save creates the Word in the database.
func (wc *WordCreate) Save(ctx context.Context) (*Word, error) {
	var (
		err  error
		node *Word
	)
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			if node, err = wc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			if wc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WordCreate) SaveX(ctx context.Context) *Word {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WordCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WordCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WordCreate) check() error {
	if _, ok := wc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Word.title"`)}
	}
	if v, ok := wc.mutation.Title(); ok {
		if err := word.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Word.title": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Meaning(); !ok {
		return &ValidationError{Name: "meaning", err: errors.New(`ent: missing required field "Word.meaning"`)}
	}
	if _, ok := wc.mutation.IsPreposition(); !ok {
		return &ValidationError{Name: "isPreposition", err: errors.New(`ent: missing required field "Word.isPreposition"`)}
	}
	if _, ok := wc.mutation.Sentence(); !ok {
		return &ValidationError{Name: "sentence", err: errors.New(`ent: missing required field "Word.sentence"`)}
	}
	if v, ok := wc.mutation.Sentence(); ok {
		if err := word.SentenceValidator(v); err != nil {
			return &ValidationError{Name: "sentence", err: fmt.Errorf(`ent: validator failed for field "Word.sentence": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "Word.duration"`)}
	}
	if v, ok := wc.mutation.Duration(); ok {
		if err := word.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "Word.duration": %w`, err)}
		}
	}
	if _, ok := wc.mutation.Start(); !ok {
		return &ValidationError{Name: "start", err: errors.New(`ent: missing required field "Word.start"`)}
	}
	if v, ok := wc.mutation.Start(); ok {
		if err := word.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "Word.start": %w`, err)}
		}
	}
	if _, ok := wc.mutation.End(); !ok {
		return &ValidationError{Name: "end", err: errors.New(`ent: missing required field "Word.end"`)}
	}
	if v, ok := wc.mutation.End(); ok {
		if err := word.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "Word.end": %w`, err)}
		}
	}
	return nil
}

func (wc *WordCreate) sqlSave(ctx context.Context) (*Word, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wc *WordCreate) createSpec() (*Word, *sqlgraph.CreateSpec) {
	var (
		_node = &Word{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: word.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: word.FieldID,
			},
		}
	)
	if value, ok := wc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := wc.mutation.Meaning(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldMeaning,
		})
		_node.Meaning = value
	}
	if value, ok := wc.mutation.IsPreposition(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: word.FieldIsPreposition,
		})
		_node.IsPreposition = value
	}
	if value, ok := wc.mutation.Sentence(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldSentence,
		})
		_node.Sentence = value
	}
	if value, ok := wc.mutation.Duration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldDuration,
		})
		_node.Duration = value
	}
	if value, ok := wc.mutation.Start(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldStart,
		})
		_node.Start = value
	}
	if value, ok := wc.mutation.End(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldEnd,
		})
		_node.End = value
	}
	if nodes := wc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   word.UserTable,
			Columns: []string{word.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_favorite_words = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   word.FileTable,
			Columns: []string{word.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: fileentity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.file_entity_words = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   word.CategoryTable,
			Columns: word.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WordCreateBulk is the builder for creating many Word entities in bulk.
type WordCreateBulk struct {
	config
	builders []*WordCreate
}

// Save creates the Word entities in the database.
func (wcb *WordCreateBulk) Save(ctx context.Context) ([]*Word, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Word, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WordMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WordCreateBulk) SaveX(ctx context.Context) []*Word {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WordCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WordCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
