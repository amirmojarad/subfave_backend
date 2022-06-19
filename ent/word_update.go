// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"interface_project/ent/collection"
	"interface_project/ent/fileentity"
	"interface_project/ent/predicate"
	"interface_project/ent/user"
	"interface_project/ent/word"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WordUpdate is the builder for updating Word entities.
type WordUpdate struct {
	config
	hooks    []Hook
	mutation *WordMutation
}

// Where appends a list predicates to the WordUpdate builder.
func (wu *WordUpdate) Where(ps ...predicate.Word) *WordUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetTitle sets the "title" field.
func (wu *WordUpdate) SetTitle(s string) *WordUpdate {
	wu.mutation.SetTitle(s)
	return wu
}

// SetMeaning sets the "meaning" field.
func (wu *WordUpdate) SetMeaning(s string) *WordUpdate {
	wu.mutation.SetMeaning(s)
	return wu
}

// SetIsPreposition sets the "isPreposition" field.
func (wu *WordUpdate) SetIsPreposition(b bool) *WordUpdate {
	wu.mutation.SetIsPreposition(b)
	return wu
}

// SetSentence sets the "sentence" field.
func (wu *WordUpdate) SetSentence(s string) *WordUpdate {
	wu.mutation.SetSentence(s)
	return wu
}

// SetDuration sets the "duration" field.
func (wu *WordUpdate) SetDuration(s string) *WordUpdate {
	wu.mutation.SetDuration(s)
	return wu
}

// SetStart sets the "start" field.
func (wu *WordUpdate) SetStart(s string) *WordUpdate {
	wu.mutation.SetStart(s)
	return wu
}

// SetEnd sets the "end" field.
func (wu *WordUpdate) SetEnd(s string) *WordUpdate {
	wu.mutation.SetEnd(s)
	return wu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (wu *WordUpdate) SetUserID(id int) *WordUpdate {
	wu.mutation.SetUserID(id)
	return wu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (wu *WordUpdate) SetNillableUserID(id *int) *WordUpdate {
	if id != nil {
		wu = wu.SetUserID(*id)
	}
	return wu
}

// SetUser sets the "user" edge to the User entity.
func (wu *WordUpdate) SetUser(u *User) *WordUpdate {
	return wu.SetUserID(u.ID)
}

// SetFileID sets the "file" edge to the FileEntity entity by ID.
func (wu *WordUpdate) SetFileID(id int) *WordUpdate {
	wu.mutation.SetFileID(id)
	return wu
}

// SetNillableFileID sets the "file" edge to the FileEntity entity by ID if the given value is not nil.
func (wu *WordUpdate) SetNillableFileID(id *int) *WordUpdate {
	if id != nil {
		wu = wu.SetFileID(*id)
	}
	return wu
}

// SetFile sets the "file" edge to the FileEntity entity.
func (wu *WordUpdate) SetFile(f *FileEntity) *WordUpdate {
	return wu.SetFileID(f.ID)
}

// AddCollectionIDs adds the "collection" edge to the Collection entity by IDs.
func (wu *WordUpdate) AddCollectionIDs(ids ...int) *WordUpdate {
	wu.mutation.AddCollectionIDs(ids...)
	return wu
}

// AddCollection adds the "collection" edges to the Collection entity.
func (wu *WordUpdate) AddCollection(c ...*Collection) *WordUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wu.AddCollectionIDs(ids...)
}

// Mutation returns the WordMutation object of the builder.
func (wu *WordUpdate) Mutation() *WordMutation {
	return wu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wu *WordUpdate) ClearUser() *WordUpdate {
	wu.mutation.ClearUser()
	return wu
}

// ClearFile clears the "file" edge to the FileEntity entity.
func (wu *WordUpdate) ClearFile() *WordUpdate {
	wu.mutation.ClearFile()
	return wu
}

// ClearCollection clears all "collection" edges to the Collection entity.
func (wu *WordUpdate) ClearCollection() *WordUpdate {
	wu.mutation.ClearCollection()
	return wu
}

// RemoveCollectionIDs removes the "collection" edge to Collection entities by IDs.
func (wu *WordUpdate) RemoveCollectionIDs(ids ...int) *WordUpdate {
	wu.mutation.RemoveCollectionIDs(ids...)
	return wu
}

// RemoveCollection removes "collection" edges to Collection entities.
func (wu *WordUpdate) RemoveCollection(c ...*Collection) *WordUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wu.RemoveCollectionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WordUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WordUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WordUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WordUpdate) check() error {
	if v, ok := wu.mutation.Title(); ok {
		if err := word.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Word.title": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Sentence(); ok {
		if err := word.SentenceValidator(v); err != nil {
			return &ValidationError{Name: "sentence", err: fmt.Errorf(`ent: validator failed for field "Word.sentence": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Duration(); ok {
		if err := word.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "Word.duration": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Start(); ok {
		if err := word.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "Word.start": %w`, err)}
		}
	}
	if v, ok := wu.mutation.End(); ok {
		if err := word.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "Word.end": %w`, err)}
		}
	}
	return nil
}

func (wu *WordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   word.Table,
			Columns: word.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: word.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldTitle,
		})
	}
	if value, ok := wu.mutation.Meaning(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldMeaning,
		})
	}
	if value, ok := wu.mutation.IsPreposition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: word.FieldIsPreposition,
		})
	}
	if value, ok := wu.mutation.Sentence(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldSentence,
		})
	}
	if value, ok := wu.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldDuration,
		})
	}
	if value, ok := wu.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldStart,
		})
	}
	if value, ok := wu.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldEnd,
		})
	}
	if wu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.FileCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.FileIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.CollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   word.CollectionTable,
			Columns: word.CollectionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: collection.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedCollectionIDs(); len(nodes) > 0 && !wu.mutation.CollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   word.CollectionTable,
			Columns: word.CollectionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: collection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.CollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   word.CollectionTable,
			Columns: word.CollectionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: collection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{word.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WordUpdateOne is the builder for updating a single Word entity.
type WordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WordMutation
}

// SetTitle sets the "title" field.
func (wuo *WordUpdateOne) SetTitle(s string) *WordUpdateOne {
	wuo.mutation.SetTitle(s)
	return wuo
}

// SetMeaning sets the "meaning" field.
func (wuo *WordUpdateOne) SetMeaning(s string) *WordUpdateOne {
	wuo.mutation.SetMeaning(s)
	return wuo
}

// SetIsPreposition sets the "isPreposition" field.
func (wuo *WordUpdateOne) SetIsPreposition(b bool) *WordUpdateOne {
	wuo.mutation.SetIsPreposition(b)
	return wuo
}

// SetSentence sets the "sentence" field.
func (wuo *WordUpdateOne) SetSentence(s string) *WordUpdateOne {
	wuo.mutation.SetSentence(s)
	return wuo
}

// SetDuration sets the "duration" field.
func (wuo *WordUpdateOne) SetDuration(s string) *WordUpdateOne {
	wuo.mutation.SetDuration(s)
	return wuo
}

// SetStart sets the "start" field.
func (wuo *WordUpdateOne) SetStart(s string) *WordUpdateOne {
	wuo.mutation.SetStart(s)
	return wuo
}

// SetEnd sets the "end" field.
func (wuo *WordUpdateOne) SetEnd(s string) *WordUpdateOne {
	wuo.mutation.SetEnd(s)
	return wuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (wuo *WordUpdateOne) SetUserID(id int) *WordUpdateOne {
	wuo.mutation.SetUserID(id)
	return wuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (wuo *WordUpdateOne) SetNillableUserID(id *int) *WordUpdateOne {
	if id != nil {
		wuo = wuo.SetUserID(*id)
	}
	return wuo
}

// SetUser sets the "user" edge to the User entity.
func (wuo *WordUpdateOne) SetUser(u *User) *WordUpdateOne {
	return wuo.SetUserID(u.ID)
}

// SetFileID sets the "file" edge to the FileEntity entity by ID.
func (wuo *WordUpdateOne) SetFileID(id int) *WordUpdateOne {
	wuo.mutation.SetFileID(id)
	return wuo
}

// SetNillableFileID sets the "file" edge to the FileEntity entity by ID if the given value is not nil.
func (wuo *WordUpdateOne) SetNillableFileID(id *int) *WordUpdateOne {
	if id != nil {
		wuo = wuo.SetFileID(*id)
	}
	return wuo
}

// SetFile sets the "file" edge to the FileEntity entity.
func (wuo *WordUpdateOne) SetFile(f *FileEntity) *WordUpdateOne {
	return wuo.SetFileID(f.ID)
}

// AddCollectionIDs adds the "collection" edge to the Collection entity by IDs.
func (wuo *WordUpdateOne) AddCollectionIDs(ids ...int) *WordUpdateOne {
	wuo.mutation.AddCollectionIDs(ids...)
	return wuo
}

// AddCollection adds the "collection" edges to the Collection entity.
func (wuo *WordUpdateOne) AddCollection(c ...*Collection) *WordUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wuo.AddCollectionIDs(ids...)
}

// Mutation returns the WordMutation object of the builder.
func (wuo *WordUpdateOne) Mutation() *WordMutation {
	return wuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wuo *WordUpdateOne) ClearUser() *WordUpdateOne {
	wuo.mutation.ClearUser()
	return wuo
}

// ClearFile clears the "file" edge to the FileEntity entity.
func (wuo *WordUpdateOne) ClearFile() *WordUpdateOne {
	wuo.mutation.ClearFile()
	return wuo
}

// ClearCollection clears all "collection" edges to the Collection entity.
func (wuo *WordUpdateOne) ClearCollection() *WordUpdateOne {
	wuo.mutation.ClearCollection()
	return wuo
}

// RemoveCollectionIDs removes the "collection" edge to Collection entities by IDs.
func (wuo *WordUpdateOne) RemoveCollectionIDs(ids ...int) *WordUpdateOne {
	wuo.mutation.RemoveCollectionIDs(ids...)
	return wuo
}

// RemoveCollection removes "collection" edges to Collection entities.
func (wuo *WordUpdateOne) RemoveCollection(c ...*Collection) *WordUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return wuo.RemoveCollectionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WordUpdateOne) Select(field string, fields ...string) *WordUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Word entity.
func (wuo *WordUpdateOne) Save(ctx context.Context) (*Word, error) {
	var (
		err  error
		node *Word
	)
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WordUpdateOne) SaveX(ctx context.Context) *Word {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WordUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WordUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WordUpdateOne) check() error {
	if v, ok := wuo.mutation.Title(); ok {
		if err := word.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Word.title": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Sentence(); ok {
		if err := word.SentenceValidator(v); err != nil {
			return &ValidationError{Name: "sentence", err: fmt.Errorf(`ent: validator failed for field "Word.sentence": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Duration(); ok {
		if err := word.DurationValidator(v); err != nil {
			return &ValidationError{Name: "duration", err: fmt.Errorf(`ent: validator failed for field "Word.duration": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Start(); ok {
		if err := word.StartValidator(v); err != nil {
			return &ValidationError{Name: "start", err: fmt.Errorf(`ent: validator failed for field "Word.start": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.End(); ok {
		if err := word.EndValidator(v); err != nil {
			return &ValidationError{Name: "end", err: fmt.Errorf(`ent: validator failed for field "Word.end": %w`, err)}
		}
	}
	return nil
}

func (wuo *WordUpdateOne) sqlSave(ctx context.Context) (_node *Word, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   word.Table,
			Columns: word.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: word.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Word.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, word.FieldID)
		for _, f := range fields {
			if !word.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != word.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldTitle,
		})
	}
	if value, ok := wuo.mutation.Meaning(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldMeaning,
		})
	}
	if value, ok := wuo.mutation.IsPreposition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: word.FieldIsPreposition,
		})
	}
	if value, ok := wuo.mutation.Sentence(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldSentence,
		})
	}
	if value, ok := wuo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldDuration,
		})
	}
	if value, ok := wuo.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldStart,
		})
	}
	if value, ok := wuo.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: word.FieldEnd,
		})
	}
	if wuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.FileCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.FileIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.CollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   word.CollectionTable,
			Columns: word.CollectionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: collection.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedCollectionIDs(); len(nodes) > 0 && !wuo.mutation.CollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   word.CollectionTable,
			Columns: word.CollectionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: collection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.CollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   word.CollectionTable,
			Columns: word.CollectionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: collection.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Word{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{word.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
