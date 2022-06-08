// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"interface_project/ent/fileentity"
	"interface_project/ent/movie"
	"interface_project/ent/predicate"
	"interface_project/ent/searchkeyword"
	"interface_project/ent/user"
	"interface_project/ent/word"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetFullName sets the "full_name" field.
func (uu *UserUpdate) SetFullName(s string) *UserUpdate {
	uu.mutation.SetFullName(s)
	return uu
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFullName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFullName(*s)
	}
	return uu
}

// ClearFullName clears the value of the "full_name" field.
func (uu *UserUpdate) ClearFullName() *UserUpdate {
	uu.mutation.ClearFullName()
	return uu
}

// SetCreatedDate sets the "created_date" field.
func (uu *UserUpdate) SetCreatedDate(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedDate(t)
	return uu
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedDate(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedDate(*t)
	}
	return uu
}

// SetUpdatedDate sets the "updated_date" field.
func (uu *UserUpdate) SetUpdatedDate(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedDate(t)
	return uu
}

// SetNillableUpdatedDate sets the "updated_date" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedDate(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdatedDate(*t)
	}
	return uu
}

// SetIsAdmin sets the "is_admin" field.
func (uu *UserUpdate) SetIsAdmin(b bool) *UserUpdate {
	uu.mutation.SetIsAdmin(b)
	return uu
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsAdmin(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsAdmin(*b)
	}
	return uu
}

// SetImageURL sets the "image_url" field.
func (uu *UserUpdate) SetImageURL(s string) *UserUpdate {
	uu.mutation.SetImageURL(s)
	return uu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillableImageURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetImageURL(*s)
	}
	return uu
}

// ClearImageURL clears the value of the "image_url" field.
func (uu *UserUpdate) ClearImageURL() *UserUpdate {
	uu.mutation.ClearImageURL()
	return uu
}

// AddFavoriteMovieIDs adds the "favorite_movies" edge to the Movie entity by IDs.
func (uu *UserUpdate) AddFavoriteMovieIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFavoriteMovieIDs(ids...)
	return uu
}

// AddFavoriteMovies adds the "favorite_movies" edges to the Movie entity.
func (uu *UserUpdate) AddFavoriteMovies(m ...*Movie) *UserUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.AddFavoriteMovieIDs(ids...)
}

// AddSearchedKeywordIDs adds the "searched_keywords" edge to the SearchKeyword entity by IDs.
func (uu *UserUpdate) AddSearchedKeywordIDs(ids ...int) *UserUpdate {
	uu.mutation.AddSearchedKeywordIDs(ids...)
	return uu
}

// AddSearchedKeywords adds the "searched_keywords" edges to the SearchKeyword entity.
func (uu *UserUpdate) AddSearchedKeywords(s ...*SearchKeyword) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddSearchedKeywordIDs(ids...)
}

// AddFavoriteWordIDs adds the "favorite_words" edge to the Word entity by IDs.
func (uu *UserUpdate) AddFavoriteWordIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFavoriteWordIDs(ids...)
	return uu
}

// AddFavoriteWords adds the "favorite_words" edges to the Word entity.
func (uu *UserUpdate) AddFavoriteWords(w ...*Word) *UserUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uu.AddFavoriteWordIDs(ids...)
}

// AddFileIDs adds the "files" edge to the FileEntity entity by IDs.
func (uu *UserUpdate) AddFileIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFileIDs(ids...)
	return uu
}

// AddFiles adds the "files" edges to the FileEntity entity.
func (uu *UserUpdate) AddFiles(f ...*FileEntity) *UserUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.AddFileIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearFavoriteMovies clears all "favorite_movies" edges to the Movie entity.
func (uu *UserUpdate) ClearFavoriteMovies() *UserUpdate {
	uu.mutation.ClearFavoriteMovies()
	return uu
}

// RemoveFavoriteMovieIDs removes the "favorite_movies" edge to Movie entities by IDs.
func (uu *UserUpdate) RemoveFavoriteMovieIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFavoriteMovieIDs(ids...)
	return uu
}

// RemoveFavoriteMovies removes "favorite_movies" edges to Movie entities.
func (uu *UserUpdate) RemoveFavoriteMovies(m ...*Movie) *UserUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uu.RemoveFavoriteMovieIDs(ids...)
}

// ClearSearchedKeywords clears all "searched_keywords" edges to the SearchKeyword entity.
func (uu *UserUpdate) ClearSearchedKeywords() *UserUpdate {
	uu.mutation.ClearSearchedKeywords()
	return uu
}

// RemoveSearchedKeywordIDs removes the "searched_keywords" edge to SearchKeyword entities by IDs.
func (uu *UserUpdate) RemoveSearchedKeywordIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveSearchedKeywordIDs(ids...)
	return uu
}

// RemoveSearchedKeywords removes "searched_keywords" edges to SearchKeyword entities.
func (uu *UserUpdate) RemoveSearchedKeywords(s ...*SearchKeyword) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveSearchedKeywordIDs(ids...)
}

// ClearFavoriteWords clears all "favorite_words" edges to the Word entity.
func (uu *UserUpdate) ClearFavoriteWords() *UserUpdate {
	uu.mutation.ClearFavoriteWords()
	return uu
}

// RemoveFavoriteWordIDs removes the "favorite_words" edge to Word entities by IDs.
func (uu *UserUpdate) RemoveFavoriteWordIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFavoriteWordIDs(ids...)
	return uu
}

// RemoveFavoriteWords removes "favorite_words" edges to Word entities.
func (uu *UserUpdate) RemoveFavoriteWords(w ...*Word) *UserUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uu.RemoveFavoriteWordIDs(ids...)
}

// ClearFiles clears all "files" edges to the FileEntity entity.
func (uu *UserUpdate) ClearFiles() *UserUpdate {
	uu.mutation.ClearFiles()
	return uu
}

// RemoveFileIDs removes the "files" edge to FileEntity entities by IDs.
func (uu *UserUpdate) RemoveFileIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFileIDs(ids...)
	return uu
}

// RemoveFiles removes "files" edges to FileEntity entities.
func (uu *UserUpdate) RemoveFiles(f ...*FileEntity) *UserUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uu.mutation.FullName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFullName,
		})
	}
	if uu.mutation.FullNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldFullName,
		})
	}
	if value, ok := uu.mutation.CreatedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedDate,
		})
	}
	if value, ok := uu.mutation.UpdatedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedDate,
		})
	}
	if value, ok := uu.mutation.IsAdmin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldIsAdmin,
		})
	}
	if value, ok := uu.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldImageURL,
		})
	}
	if uu.mutation.ImageURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldImageURL,
		})
	}
	if uu.mutation.FavoriteMoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteMoviesTable,
			Columns: user.FavoriteMoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFavoriteMoviesIDs(); len(nodes) > 0 && !uu.mutation.FavoriteMoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteMoviesTable,
			Columns: user.FavoriteMoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FavoriteMoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteMoviesTable,
			Columns: user.FavoriteMoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.SearchedKeywordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SearchedKeywordsTable,
			Columns: []string{user.SearchedKeywordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: searchkeyword.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedSearchedKeywordsIDs(); len(nodes) > 0 && !uu.mutation.SearchedKeywordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SearchedKeywordsTable,
			Columns: []string{user.SearchedKeywordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: searchkeyword.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SearchedKeywordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SearchedKeywordsTable,
			Columns: []string{user.SearchedKeywordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: searchkeyword.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FavoriteWordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoriteWordsTable,
			Columns: []string{user.FavoriteWordsColumn},
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
	if nodes := uu.mutation.RemovedFavoriteWordsIDs(); len(nodes) > 0 && !uu.mutation.FavoriteWordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoriteWordsTable,
			Columns: []string{user.FavoriteWordsColumn},
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
	if nodes := uu.mutation.FavoriteWordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoriteWordsTable,
			Columns: []string{user.FavoriteWordsColumn},
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
	if uu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
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
	if nodes := uu.mutation.RemovedFilesIDs(); len(nodes) > 0 && !uu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetFullName sets the "full_name" field.
func (uuo *UserUpdateOne) SetFullName(s string) *UserUpdateOne {
	uuo.mutation.SetFullName(s)
	return uuo
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFullName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFullName(*s)
	}
	return uuo
}

// ClearFullName clears the value of the "full_name" field.
func (uuo *UserUpdateOne) ClearFullName() *UserUpdateOne {
	uuo.mutation.ClearFullName()
	return uuo
}

// SetCreatedDate sets the "created_date" field.
func (uuo *UserUpdateOne) SetCreatedDate(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedDate(t)
	return uuo
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedDate(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedDate(*t)
	}
	return uuo
}

// SetUpdatedDate sets the "updated_date" field.
func (uuo *UserUpdateOne) SetUpdatedDate(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedDate(t)
	return uuo
}

// SetNillableUpdatedDate sets the "updated_date" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedDate(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdatedDate(*t)
	}
	return uuo
}

// SetIsAdmin sets the "is_admin" field.
func (uuo *UserUpdateOne) SetIsAdmin(b bool) *UserUpdateOne {
	uuo.mutation.SetIsAdmin(b)
	return uuo
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsAdmin(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsAdmin(*b)
	}
	return uuo
}

// SetImageURL sets the "image_url" field.
func (uuo *UserUpdateOne) SetImageURL(s string) *UserUpdateOne {
	uuo.mutation.SetImageURL(s)
	return uuo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableImageURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetImageURL(*s)
	}
	return uuo
}

// ClearImageURL clears the value of the "image_url" field.
func (uuo *UserUpdateOne) ClearImageURL() *UserUpdateOne {
	uuo.mutation.ClearImageURL()
	return uuo
}

// AddFavoriteMovieIDs adds the "favorite_movies" edge to the Movie entity by IDs.
func (uuo *UserUpdateOne) AddFavoriteMovieIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFavoriteMovieIDs(ids...)
	return uuo
}

// AddFavoriteMovies adds the "favorite_movies" edges to the Movie entity.
func (uuo *UserUpdateOne) AddFavoriteMovies(m ...*Movie) *UserUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.AddFavoriteMovieIDs(ids...)
}

// AddSearchedKeywordIDs adds the "searched_keywords" edge to the SearchKeyword entity by IDs.
func (uuo *UserUpdateOne) AddSearchedKeywordIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddSearchedKeywordIDs(ids...)
	return uuo
}

// AddSearchedKeywords adds the "searched_keywords" edges to the SearchKeyword entity.
func (uuo *UserUpdateOne) AddSearchedKeywords(s ...*SearchKeyword) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddSearchedKeywordIDs(ids...)
}

// AddFavoriteWordIDs adds the "favorite_words" edge to the Word entity by IDs.
func (uuo *UserUpdateOne) AddFavoriteWordIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFavoriteWordIDs(ids...)
	return uuo
}

// AddFavoriteWords adds the "favorite_words" edges to the Word entity.
func (uuo *UserUpdateOne) AddFavoriteWords(w ...*Word) *UserUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uuo.AddFavoriteWordIDs(ids...)
}

// AddFileIDs adds the "files" edge to the FileEntity entity by IDs.
func (uuo *UserUpdateOne) AddFileIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFileIDs(ids...)
	return uuo
}

// AddFiles adds the "files" edges to the FileEntity entity.
func (uuo *UserUpdateOne) AddFiles(f ...*FileEntity) *UserUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.AddFileIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearFavoriteMovies clears all "favorite_movies" edges to the Movie entity.
func (uuo *UserUpdateOne) ClearFavoriteMovies() *UserUpdateOne {
	uuo.mutation.ClearFavoriteMovies()
	return uuo
}

// RemoveFavoriteMovieIDs removes the "favorite_movies" edge to Movie entities by IDs.
func (uuo *UserUpdateOne) RemoveFavoriteMovieIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFavoriteMovieIDs(ids...)
	return uuo
}

// RemoveFavoriteMovies removes "favorite_movies" edges to Movie entities.
func (uuo *UserUpdateOne) RemoveFavoriteMovies(m ...*Movie) *UserUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uuo.RemoveFavoriteMovieIDs(ids...)
}

// ClearSearchedKeywords clears all "searched_keywords" edges to the SearchKeyword entity.
func (uuo *UserUpdateOne) ClearSearchedKeywords() *UserUpdateOne {
	uuo.mutation.ClearSearchedKeywords()
	return uuo
}

// RemoveSearchedKeywordIDs removes the "searched_keywords" edge to SearchKeyword entities by IDs.
func (uuo *UserUpdateOne) RemoveSearchedKeywordIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveSearchedKeywordIDs(ids...)
	return uuo
}

// RemoveSearchedKeywords removes "searched_keywords" edges to SearchKeyword entities.
func (uuo *UserUpdateOne) RemoveSearchedKeywords(s ...*SearchKeyword) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveSearchedKeywordIDs(ids...)
}

// ClearFavoriteWords clears all "favorite_words" edges to the Word entity.
func (uuo *UserUpdateOne) ClearFavoriteWords() *UserUpdateOne {
	uuo.mutation.ClearFavoriteWords()
	return uuo
}

// RemoveFavoriteWordIDs removes the "favorite_words" edge to Word entities by IDs.
func (uuo *UserUpdateOne) RemoveFavoriteWordIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFavoriteWordIDs(ids...)
	return uuo
}

// RemoveFavoriteWords removes "favorite_words" edges to Word entities.
func (uuo *UserUpdateOne) RemoveFavoriteWords(w ...*Word) *UserUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uuo.RemoveFavoriteWordIDs(ids...)
}

// ClearFiles clears all "files" edges to the FileEntity entity.
func (uuo *UserUpdateOne) ClearFiles() *UserUpdateOne {
	uuo.mutation.ClearFiles()
	return uuo
}

// RemoveFileIDs removes the "files" edge to FileEntity entities by IDs.
func (uuo *UserUpdateOne) RemoveFileIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFileIDs(ids...)
	return uuo
}

// RemoveFiles removes "files" edges to FileEntity entities.
func (uuo *UserUpdateOne) RemoveFiles(f ...*FileEntity) *UserUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.RemoveFileIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
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
	if value, ok := uuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if value, ok := uuo.mutation.FullName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFullName,
		})
	}
	if uuo.mutation.FullNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldFullName,
		})
	}
	if value, ok := uuo.mutation.CreatedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldCreatedDate,
		})
	}
	if value, ok := uuo.mutation.UpdatedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedDate,
		})
	}
	if value, ok := uuo.mutation.IsAdmin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldIsAdmin,
		})
	}
	if value, ok := uuo.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldImageURL,
		})
	}
	if uuo.mutation.ImageURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: user.FieldImageURL,
		})
	}
	if uuo.mutation.FavoriteMoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteMoviesTable,
			Columns: user.FavoriteMoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFavoriteMoviesIDs(); len(nodes) > 0 && !uuo.mutation.FavoriteMoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteMoviesTable,
			Columns: user.FavoriteMoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FavoriteMoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FavoriteMoviesTable,
			Columns: user.FavoriteMoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.SearchedKeywordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SearchedKeywordsTable,
			Columns: []string{user.SearchedKeywordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: searchkeyword.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedSearchedKeywordsIDs(); len(nodes) > 0 && !uuo.mutation.SearchedKeywordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SearchedKeywordsTable,
			Columns: []string{user.SearchedKeywordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: searchkeyword.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SearchedKeywordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SearchedKeywordsTable,
			Columns: []string{user.SearchedKeywordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: searchkeyword.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FavoriteWordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoriteWordsTable,
			Columns: []string{user.FavoriteWordsColumn},
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
	if nodes := uuo.mutation.RemovedFavoriteWordsIDs(); len(nodes) > 0 && !uuo.mutation.FavoriteWordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoriteWordsTable,
			Columns: []string{user.FavoriteWordsColumn},
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
	if nodes := uuo.mutation.FavoriteWordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FavoriteWordsTable,
			Columns: []string{user.FavoriteWordsColumn},
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
	if uuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
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
	if nodes := uuo.mutation.RemovedFilesIDs(); len(nodes) > 0 && !uuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.FilesTable,
			Columns: []string{user.FilesColumn},
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
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
