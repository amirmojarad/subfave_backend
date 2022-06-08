// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"interface_project/ent/fileentity"
	"interface_project/ent/predicate"
	"interface_project/ent/word"
	"interface_project/ent/wordnode"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WordNodeUpdate is the builder for updating WordNode entities.
type WordNodeUpdate struct {
	config
	hooks    []Hook
	mutation *WordNodeMutation
}

// Where appends a list predicates to the WordNodeUpdate builder.
func (wnu *WordNodeUpdate) Where(ps ...predicate.WordNode) *WordNodeUpdate {
	wnu.mutation.Where(ps...)
	return wnu
}

// SetTitle sets the "title" field.
func (wnu *WordNodeUpdate) SetTitle(s string) *WordNodeUpdate {
	wnu.mutation.SetTitle(s)
	return wnu
}

// SetIsPreposition sets the "is_preposition" field.
func (wnu *WordNodeUpdate) SetIsPreposition(b bool) *WordNodeUpdate {
	wnu.mutation.SetIsPreposition(b)
	return wnu
}

// SetOccurence sets the "occurence" field.
func (wnu *WordNodeUpdate) SetOccurence(i int) *WordNodeUpdate {
	wnu.mutation.ResetOccurence()
	wnu.mutation.SetOccurence(i)
	return wnu
}

// SetNillableOccurence sets the "occurence" field if the given value is not nil.
func (wnu *WordNodeUpdate) SetNillableOccurence(i *int) *WordNodeUpdate {
	if i != nil {
		wnu.SetOccurence(*i)
	}
	return wnu
}

// AddOccurence adds i to the "occurence" field.
func (wnu *WordNodeUpdate) AddOccurence(i int) *WordNodeUpdate {
	wnu.mutation.AddOccurence(i)
	return wnu
}

// ClearOccurence clears the value of the "occurence" field.
func (wnu *WordNodeUpdate) ClearOccurence() *WordNodeUpdate {
	wnu.mutation.ClearOccurence()
	return wnu
}

// AddWordIDs adds the "words" edge to the Word entity by IDs.
func (wnu *WordNodeUpdate) AddWordIDs(ids ...int) *WordNodeUpdate {
	wnu.mutation.AddWordIDs(ids...)
	return wnu
}

// AddWords adds the "words" edges to the Word entity.
func (wnu *WordNodeUpdate) AddWords(w ...*Word) *WordNodeUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wnu.AddWordIDs(ids...)
}

// SetFileID sets the "file" edge to the FileEntity entity by ID.
func (wnu *WordNodeUpdate) SetFileID(id int) *WordNodeUpdate {
	wnu.mutation.SetFileID(id)
	return wnu
}

// SetFile sets the "file" edge to the FileEntity entity.
func (wnu *WordNodeUpdate) SetFile(f *FileEntity) *WordNodeUpdate {
	return wnu.SetFileID(f.ID)
}

// Mutation returns the WordNodeMutation object of the builder.
func (wnu *WordNodeUpdate) Mutation() *WordNodeMutation {
	return wnu.mutation
}

// ClearWords clears all "words" edges to the Word entity.
func (wnu *WordNodeUpdate) ClearWords() *WordNodeUpdate {
	wnu.mutation.ClearWords()
	return wnu
}

// RemoveWordIDs removes the "words" edge to Word entities by IDs.
func (wnu *WordNodeUpdate) RemoveWordIDs(ids ...int) *WordNodeUpdate {
	wnu.mutation.RemoveWordIDs(ids...)
	return wnu
}

// RemoveWords removes "words" edges to Word entities.
func (wnu *WordNodeUpdate) RemoveWords(w ...*Word) *WordNodeUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wnu.RemoveWordIDs(ids...)
}

// ClearFile clears the "file" edge to the FileEntity entity.
func (wnu *WordNodeUpdate) ClearFile() *WordNodeUpdate {
	wnu.mutation.ClearFile()
	return wnu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wnu *WordNodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wnu.hooks) == 0 {
		if err = wnu.check(); err != nil {
			return 0, err
		}
		affected, err = wnu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WordNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wnu.check(); err != nil {
				return 0, err
			}
			wnu.mutation = mutation
			affected, err = wnu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wnu.hooks) - 1; i >= 0; i-- {
			if wnu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wnu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wnu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wnu *WordNodeUpdate) SaveX(ctx context.Context) int {
	affected, err := wnu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wnu *WordNodeUpdate) Exec(ctx context.Context) error {
	_, err := wnu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wnu *WordNodeUpdate) ExecX(ctx context.Context) {
	if err := wnu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wnu *WordNodeUpdate) check() error {
	if v, ok := wnu.mutation.Title(); ok {
		if err := wordnode.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "WordNode.title": %w`, err)}
		}
	}
	if _, ok := wnu.mutation.FileID(); wnu.mutation.FileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WordNode.file"`)
	}
	return nil
}

func (wnu *WordNodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wordnode.Table,
			Columns: wordnode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wordnode.FieldID,
			},
		},
	}
	if ps := wnu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wnu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wordnode.FieldTitle,
		})
	}
	if value, ok := wnu.mutation.IsPreposition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wordnode.FieldIsPreposition,
		})
	}
	if value, ok := wnu.mutation.Occurence(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wordnode.FieldOccurence,
		})
	}
	if value, ok := wnu.mutation.AddedOccurence(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wordnode.FieldOccurence,
		})
	}
	if wnu.mutation.OccurenceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: wordnode.FieldOccurence,
		})
	}
	if wnu.mutation.WordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wordnode.WordsTable,
			Columns: []string{wordnode.WordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: word.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wnu.mutation.RemovedWordsIDs(); len(nodes) > 0 && !wnu.mutation.WordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wordnode.WordsTable,
			Columns: []string{wordnode.WordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: word.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wnu.mutation.WordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wordnode.WordsTable,
			Columns: []string{wordnode.WordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: word.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wnu.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wordnode.FileTable,
			Columns: []string{wordnode.FileColumn},
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
	if nodes := wnu.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wordnode.FileTable,
			Columns: []string{wordnode.FileColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, wnu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wordnode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WordNodeUpdateOne is the builder for updating a single WordNode entity.
type WordNodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WordNodeMutation
}

// SetTitle sets the "title" field.
func (wnuo *WordNodeUpdateOne) SetTitle(s string) *WordNodeUpdateOne {
	wnuo.mutation.SetTitle(s)
	return wnuo
}

// SetIsPreposition sets the "is_preposition" field.
func (wnuo *WordNodeUpdateOne) SetIsPreposition(b bool) *WordNodeUpdateOne {
	wnuo.mutation.SetIsPreposition(b)
	return wnuo
}

// SetOccurence sets the "occurence" field.
func (wnuo *WordNodeUpdateOne) SetOccurence(i int) *WordNodeUpdateOne {
	wnuo.mutation.ResetOccurence()
	wnuo.mutation.SetOccurence(i)
	return wnuo
}

// SetNillableOccurence sets the "occurence" field if the given value is not nil.
func (wnuo *WordNodeUpdateOne) SetNillableOccurence(i *int) *WordNodeUpdateOne {
	if i != nil {
		wnuo.SetOccurence(*i)
	}
	return wnuo
}

// AddOccurence adds i to the "occurence" field.
func (wnuo *WordNodeUpdateOne) AddOccurence(i int) *WordNodeUpdateOne {
	wnuo.mutation.AddOccurence(i)
	return wnuo
}

// ClearOccurence clears the value of the "occurence" field.
func (wnuo *WordNodeUpdateOne) ClearOccurence() *WordNodeUpdateOne {
	wnuo.mutation.ClearOccurence()
	return wnuo
}

// AddWordIDs adds the "words" edge to the Word entity by IDs.
func (wnuo *WordNodeUpdateOne) AddWordIDs(ids ...int) *WordNodeUpdateOne {
	wnuo.mutation.AddWordIDs(ids...)
	return wnuo
}

// AddWords adds the "words" edges to the Word entity.
func (wnuo *WordNodeUpdateOne) AddWords(w ...*Word) *WordNodeUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wnuo.AddWordIDs(ids...)
}

// SetFileID sets the "file" edge to the FileEntity entity by ID.
func (wnuo *WordNodeUpdateOne) SetFileID(id int) *WordNodeUpdateOne {
	wnuo.mutation.SetFileID(id)
	return wnuo
}

// SetFile sets the "file" edge to the FileEntity entity.
func (wnuo *WordNodeUpdateOne) SetFile(f *FileEntity) *WordNodeUpdateOne {
	return wnuo.SetFileID(f.ID)
}

// Mutation returns the WordNodeMutation object of the builder.
func (wnuo *WordNodeUpdateOne) Mutation() *WordNodeMutation {
	return wnuo.mutation
}

// ClearWords clears all "words" edges to the Word entity.
func (wnuo *WordNodeUpdateOne) ClearWords() *WordNodeUpdateOne {
	wnuo.mutation.ClearWords()
	return wnuo
}

// RemoveWordIDs removes the "words" edge to Word entities by IDs.
func (wnuo *WordNodeUpdateOne) RemoveWordIDs(ids ...int) *WordNodeUpdateOne {
	wnuo.mutation.RemoveWordIDs(ids...)
	return wnuo
}

// RemoveWords removes "words" edges to Word entities.
func (wnuo *WordNodeUpdateOne) RemoveWords(w ...*Word) *WordNodeUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wnuo.RemoveWordIDs(ids...)
}

// ClearFile clears the "file" edge to the FileEntity entity.
func (wnuo *WordNodeUpdateOne) ClearFile() *WordNodeUpdateOne {
	wnuo.mutation.ClearFile()
	return wnuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wnuo *WordNodeUpdateOne) Select(field string, fields ...string) *WordNodeUpdateOne {
	wnuo.fields = append([]string{field}, fields...)
	return wnuo
}

// Save executes the query and returns the updated WordNode entity.
func (wnuo *WordNodeUpdateOne) Save(ctx context.Context) (*WordNode, error) {
	var (
		err  error
		node *WordNode
	)
	if len(wnuo.hooks) == 0 {
		if err = wnuo.check(); err != nil {
			return nil, err
		}
		node, err = wnuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WordNodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wnuo.check(); err != nil {
				return nil, err
			}
			wnuo.mutation = mutation
			node, err = wnuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wnuo.hooks) - 1; i >= 0; i-- {
			if wnuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wnuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wnuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wnuo *WordNodeUpdateOne) SaveX(ctx context.Context) *WordNode {
	node, err := wnuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wnuo *WordNodeUpdateOne) Exec(ctx context.Context) error {
	_, err := wnuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wnuo *WordNodeUpdateOne) ExecX(ctx context.Context) {
	if err := wnuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wnuo *WordNodeUpdateOne) check() error {
	if v, ok := wnuo.mutation.Title(); ok {
		if err := wordnode.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "WordNode.title": %w`, err)}
		}
	}
	if _, ok := wnuo.mutation.FileID(); wnuo.mutation.FileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WordNode.file"`)
	}
	return nil
}

func (wnuo *WordNodeUpdateOne) sqlSave(ctx context.Context) (_node *WordNode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wordnode.Table,
			Columns: wordnode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wordnode.FieldID,
			},
		},
	}
	id, ok := wnuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WordNode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wnuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wordnode.FieldID)
		for _, f := range fields {
			if !wordnode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wordnode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wnuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wnuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wordnode.FieldTitle,
		})
	}
	if value, ok := wnuo.mutation.IsPreposition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wordnode.FieldIsPreposition,
		})
	}
	if value, ok := wnuo.mutation.Occurence(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wordnode.FieldOccurence,
		})
	}
	if value, ok := wnuo.mutation.AddedOccurence(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: wordnode.FieldOccurence,
		})
	}
	if wnuo.mutation.OccurenceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: wordnode.FieldOccurence,
		})
	}
	if wnuo.mutation.WordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wordnode.WordsTable,
			Columns: []string{wordnode.WordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: word.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wnuo.mutation.RemovedWordsIDs(); len(nodes) > 0 && !wnuo.mutation.WordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wordnode.WordsTable,
			Columns: []string{wordnode.WordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: word.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wnuo.mutation.WordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wordnode.WordsTable,
			Columns: []string{wordnode.WordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: word.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wnuo.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wordnode.FileTable,
			Columns: []string{wordnode.FileColumn},
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
	if nodes := wnuo.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wordnode.FileTable,
			Columns: []string{wordnode.FileColumn},
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
	_node = &WordNode{config: wnuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wnuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wordnode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
