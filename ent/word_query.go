// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"interface_project/ent/collection"
	"interface_project/ent/fileentity"
	"interface_project/ent/predicate"
	"interface_project/ent/user"
	"interface_project/ent/word"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WordQuery is the builder for querying Word entities.
type WordQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Word
	// eager-loading edges.
	withUser       *UserQuery
	withFile       *FileEntityQuery
	withCollection *CollectionQuery
	withOwner      *UserQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WordQuery builder.
func (wq *WordQuery) Where(ps ...predicate.Word) *WordQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit adds a limit step to the query.
func (wq *WordQuery) Limit(limit int) *WordQuery {
	wq.limit = &limit
	return wq
}

// Offset adds an offset step to the query.
func (wq *WordQuery) Offset(offset int) *WordQuery {
	wq.offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WordQuery) Unique(unique bool) *WordQuery {
	wq.unique = &unique
	return wq
}

// Order adds an order step to the query.
func (wq *WordQuery) Order(o ...OrderFunc) *WordQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryUser chains the current query on the "user" edge.
func (wq *WordQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, word.UserTable, word.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFile chains the current query on the "file" edge.
func (wq *WordQuery) QueryFile() *FileEntityQuery {
	query := &FileEntityQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, selector),
			sqlgraph.To(fileentity.Table, fileentity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, word.FileTable, word.FileColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCollection chains the current query on the "collection" edge.
func (wq *WordQuery) QueryCollection() *CollectionQuery {
	query := &CollectionQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, selector),
			sqlgraph.To(collection.Table, collection.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, word.CollectionTable, word.CollectionPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwner chains the current query on the "owner" edge.
func (wq *WordQuery) QueryOwner() *UserQuery {
	query := &UserQuery{config: wq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, word.OwnerTable, word.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Word entity from the query.
// Returns a *NotFoundError when no Word was found.
func (wq *WordQuery) First(ctx context.Context) (*Word, error) {
	nodes, err := wq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{word.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WordQuery) FirstX(ctx context.Context) *Word {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Word ID from the query.
// Returns a *NotFoundError when no Word ID was found.
func (wq *WordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{word.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WordQuery) FirstIDX(ctx context.Context) int {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Word entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Word entity is found.
// Returns a *NotFoundError when no Word entities are found.
func (wq *WordQuery) Only(ctx context.Context) (*Word, error) {
	nodes, err := wq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{word.Label}
	default:
		return nil, &NotSingularError{word.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WordQuery) OnlyX(ctx context.Context) *Word {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Word ID in the query.
// Returns a *NotSingularError when more than one Word ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = &NotSingularError{word.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WordQuery) OnlyIDX(ctx context.Context) int {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Words.
func (wq *WordQuery) All(ctx context.Context) ([]*Word, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wq *WordQuery) AllX(ctx context.Context) []*Word {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Word IDs.
func (wq *WordQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wq.Select(word.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WordQuery) IDsX(ctx context.Context) []int {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WordQuery) Count(ctx context.Context) (int, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WordQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WordQuery) Exist(ctx context.Context) (bool, error) {
	if err := wq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WordQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WordQuery) Clone() *WordQuery {
	if wq == nil {
		return nil
	}
	return &WordQuery{
		config:         wq.config,
		limit:          wq.limit,
		offset:         wq.offset,
		order:          append([]OrderFunc{}, wq.order...),
		predicates:     append([]predicate.Word{}, wq.predicates...),
		withUser:       wq.withUser.Clone(),
		withFile:       wq.withFile.Clone(),
		withCollection: wq.withCollection.Clone(),
		withOwner:      wq.withOwner.Clone(),
		// clone intermediate query.
		sql:    wq.sql.Clone(),
		path:   wq.path,
		unique: wq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WordQuery) WithUser(opts ...func(*UserQuery)) *WordQuery {
	query := &UserQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withUser = query
	return wq
}

// WithFile tells the query-builder to eager-load the nodes that are connected to
// the "file" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WordQuery) WithFile(opts ...func(*FileEntityQuery)) *WordQuery {
	query := &FileEntityQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withFile = query
	return wq
}

// WithCollection tells the query-builder to eager-load the nodes that are connected to
// the "collection" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WordQuery) WithCollection(opts ...func(*CollectionQuery)) *WordQuery {
	query := &CollectionQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withCollection = query
	return wq
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WordQuery) WithOwner(opts ...func(*UserQuery)) *WordQuery {
	query := &UserQuery{config: wq.config}
	for _, opt := range opts {
		opt(query)
	}
	wq.withOwner = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Word.Query().
//		GroupBy(word.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wq *WordQuery) GroupBy(field string, fields ...string) *WordGroupBy {
	group := &WordGroupBy{config: wq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Word.Query().
//		Select(word.FieldTitle).
//		Scan(ctx, &v)
//
func (wq *WordQuery) Select(fields ...string) *WordSelect {
	wq.fields = append(wq.fields, fields...)
	return &WordSelect{WordQuery: wq}
}

func (wq *WordQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wq.fields {
		if !word.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WordQuery) sqlAll(ctx context.Context) ([]*Word, error) {
	var (
		nodes       = []*Word{}
		withFKs     = wq.withFKs
		_spec       = wq.querySpec()
		loadedTypes = [4]bool{
			wq.withUser != nil,
			wq.withFile != nil,
			wq.withCollection != nil,
			wq.withOwner != nil,
		}
	)
	if wq.withUser != nil || wq.withFile != nil || wq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, word.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Word{config: wq.config}
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
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Word)
		for i := range nodes {
			if nodes[i].user_favorite_words == nil {
				continue
			}
			fk := *nodes[i].user_favorite_words
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_favorite_words" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := wq.withFile; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Word)
		for i := range nodes {
			if nodes[i].file_entity_words == nil {
				continue
			}
			fk := *nodes[i].file_entity_words
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(fileentity.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "file_entity_words" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.File = n
			}
		}
	}

	if query := wq.withCollection; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Word, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Collection = []*Collection{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Word)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   word.CollectionTable,
				Columns: word.CollectionPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(word.CollectionPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, wq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "collection": %w`, err)
		}
		query.Where(collection.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "collection" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Collection = append(nodes[i].Edges.Collection, n)
			}
		}
	}

	if query := wq.withOwner; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Word)
		for i := range nodes {
			if nodes[i].user_words == nil {
				continue
			}
			fk := *nodes[i].user_words
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
				return nil, fmt.Errorf(`unexpected foreign-key "user_words" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	return nodes, nil
}

func (wq *WordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	_spec.Node.Columns = wq.fields
	if len(wq.fields) > 0 {
		_spec.Unique = wq.unique != nil && *wq.unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wq *WordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   word.Table,
			Columns: word.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: word.FieldID,
			},
		},
		From:   wq.sql,
		Unique: true,
	}
	if unique := wq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, word.FieldID)
		for i := range fields {
			if fields[i] != word.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(word.Table)
	columns := wq.fields
	if len(columns) == 0 {
		columns = word.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.unique != nil && *wq.unique {
		selector.Distinct()
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WordGroupBy is the group-by builder for Word entities.
type WordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WordGroupBy) Aggregate(fns ...AggregateFunc) *WordGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the group-by query and scans the result into the given value.
func (wgb *WordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wgb.path(ctx)
	if err != nil {
		return err
	}
	wgb.sql = query
	return wgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wgb *WordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *WordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wgb *WordGroupBy) StringsX(ctx context.Context) []string {
	v, err := wgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *WordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wgb *WordGroupBy) StringX(ctx context.Context) string {
	v, err := wgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *WordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wgb *WordGroupBy) IntsX(ctx context.Context) []int {
	v, err := wgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *WordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wgb *WordGroupBy) IntX(ctx context.Context) int {
	v, err := wgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *WordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wgb *WordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *WordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wgb *WordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wgb *WordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wgb.fields) > 1 {
		return nil, errors.New("ent: WordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wgb *WordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wgb *WordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wgb *WordGroupBy) BoolX(ctx context.Context) bool {
	v, err := wgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wgb *WordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wgb.fields {
		if !word.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wgb *WordGroupBy) sqlQuery() *sql.Selector {
	selector := wgb.sql.Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wgb.fields)+len(wgb.fns))
		for _, f := range wgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wgb.fields...)...)
}

// WordSelect is the builder for selecting fields of Word entities.
type WordSelect struct {
	*WordQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WordSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	ws.sql = ws.WordQuery.sqlQuery(ctx)
	return ws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ws *WordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ws *WordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ws *WordSelect) StringsX(ctx context.Context) []string {
	v, err := ws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ws *WordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ws *WordSelect) StringX(ctx context.Context) string {
	v, err := ws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ws *WordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ws *WordSelect) IntsX(ctx context.Context) []int {
	v, err := ws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ws *WordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ws *WordSelect) IntX(ctx context.Context) int {
	v, err := ws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ws *WordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ws *WordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ws *WordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ws *WordSelect) Float64X(ctx context.Context) float64 {
	v, err := ws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ws *WordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ws.fields) > 1 {
		return nil, errors.New("ent: WordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ws *WordSelect) BoolsX(ctx context.Context) []bool {
	v, err := ws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ws *WordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = fmt.Errorf("ent: WordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ws *WordSelect) BoolX(ctx context.Context) bool {
	v, err := ws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ws *WordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ws.sql.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
