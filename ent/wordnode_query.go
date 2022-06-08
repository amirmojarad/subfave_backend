// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"interface_project/ent/fileentity"
	"interface_project/ent/predicate"
	"interface_project/ent/word"
	"interface_project/ent/wordnode"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WordNodeQuery is the builder for querying WordNode entities.
type WordNodeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.WordNode
	// eager-loading edges.
	withWords *WordQuery
	withFile  *FileEntityQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WordNodeQuery builder.
func (wnq *WordNodeQuery) Where(ps ...predicate.WordNode) *WordNodeQuery {
	wnq.predicates = append(wnq.predicates, ps...)
	return wnq
}

// Limit adds a limit step to the query.
func (wnq *WordNodeQuery) Limit(limit int) *WordNodeQuery {
	wnq.limit = &limit
	return wnq
}

// Offset adds an offset step to the query.
func (wnq *WordNodeQuery) Offset(offset int) *WordNodeQuery {
	wnq.offset = &offset
	return wnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wnq *WordNodeQuery) Unique(unique bool) *WordNodeQuery {
	wnq.unique = &unique
	return wnq
}

// Order adds an order step to the query.
func (wnq *WordNodeQuery) Order(o ...OrderFunc) *WordNodeQuery {
	wnq.order = append(wnq.order, o...)
	return wnq
}

// QueryWords chains the current query on the "words" edge.
func (wnq *WordNodeQuery) QueryWords() *WordQuery {
	query := &WordQuery{config: wnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(wordnode.Table, wordnode.FieldID, selector),
			sqlgraph.To(word.Table, word.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, wordnode.WordsTable, wordnode.WordsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFile chains the current query on the "file" edge.
func (wnq *WordNodeQuery) QueryFile() *FileEntityQuery {
	query := &FileEntityQuery{config: wnq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wnq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(wordnode.Table, wordnode.FieldID, selector),
			sqlgraph.To(fileentity.Table, fileentity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, wordnode.FileTable, wordnode.FileColumn),
		)
		fromU = sqlgraph.SetNeighbors(wnq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WordNode entity from the query.
// Returns a *NotFoundError when no WordNode was found.
func (wnq *WordNodeQuery) First(ctx context.Context) (*WordNode, error) {
	nodes, err := wnq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wordnode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wnq *WordNodeQuery) FirstX(ctx context.Context) *WordNode {
	node, err := wnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WordNode ID from the query.
// Returns a *NotFoundError when no WordNode ID was found.
func (wnq *WordNodeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wnq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wordnode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wnq *WordNodeQuery) FirstIDX(ctx context.Context) int {
	id, err := wnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WordNode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WordNode entity is found.
// Returns a *NotFoundError when no WordNode entities are found.
func (wnq *WordNodeQuery) Only(ctx context.Context) (*WordNode, error) {
	nodes, err := wnq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wordnode.Label}
	default:
		return nil, &NotSingularError{wordnode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wnq *WordNodeQuery) OnlyX(ctx context.Context) *WordNode {
	node, err := wnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WordNode ID in the query.
// Returns a *NotSingularError when more than one WordNode ID is found.
// Returns a *NotFoundError when no entities are found.
func (wnq *WordNodeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wnq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wordnode.Label}
	default:
		err = &NotSingularError{wordnode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wnq *WordNodeQuery) OnlyIDX(ctx context.Context) int {
	id, err := wnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WordNodes.
func (wnq *WordNodeQuery) All(ctx context.Context) ([]*WordNode, error) {
	if err := wnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wnq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wnq *WordNodeQuery) AllX(ctx context.Context) []*WordNode {
	nodes, err := wnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WordNode IDs.
func (wnq *WordNodeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wnq.Select(wordnode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wnq *WordNodeQuery) IDsX(ctx context.Context) []int {
	ids, err := wnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wnq *WordNodeQuery) Count(ctx context.Context) (int, error) {
	if err := wnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wnq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wnq *WordNodeQuery) CountX(ctx context.Context) int {
	count, err := wnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wnq *WordNodeQuery) Exist(ctx context.Context) (bool, error) {
	if err := wnq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wnq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wnq *WordNodeQuery) ExistX(ctx context.Context) bool {
	exist, err := wnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WordNodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wnq *WordNodeQuery) Clone() *WordNodeQuery {
	if wnq == nil {
		return nil
	}
	return &WordNodeQuery{
		config:     wnq.config,
		limit:      wnq.limit,
		offset:     wnq.offset,
		order:      append([]OrderFunc{}, wnq.order...),
		predicates: append([]predicate.WordNode{}, wnq.predicates...),
		withWords:  wnq.withWords.Clone(),
		withFile:   wnq.withFile.Clone(),
		// clone intermediate query.
		sql:    wnq.sql.Clone(),
		path:   wnq.path,
		unique: wnq.unique,
	}
}

// WithWords tells the query-builder to eager-load the nodes that are connected to
// the "words" edge. The optional arguments are used to configure the query builder of the edge.
func (wnq *WordNodeQuery) WithWords(opts ...func(*WordQuery)) *WordNodeQuery {
	query := &WordQuery{config: wnq.config}
	for _, opt := range opts {
		opt(query)
	}
	wnq.withWords = query
	return wnq
}

// WithFile tells the query-builder to eager-load the nodes that are connected to
// the "file" edge. The optional arguments are used to configure the query builder of the edge.
func (wnq *WordNodeQuery) WithFile(opts ...func(*FileEntityQuery)) *WordNodeQuery {
	query := &FileEntityQuery{config: wnq.config}
	for _, opt := range opts {
		opt(query)
	}
	wnq.withFile = query
	return wnq
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
//	client.WordNode.Query().
//		GroupBy(wordnode.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wnq *WordNodeQuery) GroupBy(field string, fields ...string) *WordNodeGroupBy {
	group := &WordNodeGroupBy{config: wnq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wnq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wnq.sqlQuery(ctx), nil
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
//	client.WordNode.Query().
//		Select(wordnode.FieldTitle).
//		Scan(ctx, &v)
//
func (wnq *WordNodeQuery) Select(fields ...string) *WordNodeSelect {
	wnq.fields = append(wnq.fields, fields...)
	return &WordNodeSelect{WordNodeQuery: wnq}
}

func (wnq *WordNodeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range wnq.fields {
		if !wordnode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wnq.path != nil {
		prev, err := wnq.path(ctx)
		if err != nil {
			return err
		}
		wnq.sql = prev
	}
	return nil
}

func (wnq *WordNodeQuery) sqlAll(ctx context.Context) ([]*WordNode, error) {
	var (
		nodes       = []*WordNode{}
		withFKs     = wnq.withFKs
		_spec       = wnq.querySpec()
		loadedTypes = [2]bool{
			wnq.withWords != nil,
			wnq.withFile != nil,
		}
	)
	if wnq.withFile != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, wordnode.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &WordNode{config: wnq.config}
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
	if err := sqlgraph.QueryNodes(ctx, wnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wnq.withWords; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*WordNode)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Words = []*Word{}
		}
		query.withFKs = true
		query.Where(predicate.Word(func(s *sql.Selector) {
			s.Where(sql.InValues(wordnode.WordsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.word_node_words
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "word_node_words" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "word_node_words" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Words = append(node.Edges.Words, n)
		}
	}

	if query := wnq.withFile; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*WordNode)
		for i := range nodes {
			if nodes[i].file_entity_wordnodes == nil {
				continue
			}
			fk := *nodes[i].file_entity_wordnodes
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
				return nil, fmt.Errorf(`unexpected foreign-key "file_entity_wordnodes" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.File = n
			}
		}
	}

	return nodes, nil
}

func (wnq *WordNodeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wnq.querySpec()
	_spec.Node.Columns = wnq.fields
	if len(wnq.fields) > 0 {
		_spec.Unique = wnq.unique != nil && *wnq.unique
	}
	return sqlgraph.CountNodes(ctx, wnq.driver, _spec)
}

func (wnq *WordNodeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wnq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (wnq *WordNodeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wordnode.Table,
			Columns: wordnode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: wordnode.FieldID,
			},
		},
		From:   wnq.sql,
		Unique: true,
	}
	if unique := wnq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wnq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wordnode.FieldID)
		for i := range fields {
			if fields[i] != wordnode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wnq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wnq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wnq *WordNodeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wnq.driver.Dialect())
	t1 := builder.Table(wordnode.Table)
	columns := wnq.fields
	if len(columns) == 0 {
		columns = wordnode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wnq.sql != nil {
		selector = wnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wnq.unique != nil && *wnq.unique {
		selector.Distinct()
	}
	for _, p := range wnq.predicates {
		p(selector)
	}
	for _, p := range wnq.order {
		p(selector)
	}
	if offset := wnq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wnq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WordNodeGroupBy is the group-by builder for WordNode entities.
type WordNodeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wngb *WordNodeGroupBy) Aggregate(fns ...AggregateFunc) *WordNodeGroupBy {
	wngb.fns = append(wngb.fns, fns...)
	return wngb
}

// Scan applies the group-by query and scans the result into the given value.
func (wngb *WordNodeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wngb.path(ctx)
	if err != nil {
		return err
	}
	wngb.sql = query
	return wngb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wngb *WordNodeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wngb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (wngb *WordNodeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wngb.fields) > 1 {
		return nil, errors.New("ent: WordNodeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wngb *WordNodeGroupBy) StringsX(ctx context.Context) []string {
	v, err := wngb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wngb *WordNodeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wngb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wordnode.Label}
	default:
		err = fmt.Errorf("ent: WordNodeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wngb *WordNodeGroupBy) StringX(ctx context.Context) string {
	v, err := wngb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (wngb *WordNodeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wngb.fields) > 1 {
		return nil, errors.New("ent: WordNodeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wngb *WordNodeGroupBy) IntsX(ctx context.Context) []int {
	v, err := wngb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wngb *WordNodeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wngb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wordnode.Label}
	default:
		err = fmt.Errorf("ent: WordNodeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wngb *WordNodeGroupBy) IntX(ctx context.Context) int {
	v, err := wngb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (wngb *WordNodeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wngb.fields) > 1 {
		return nil, errors.New("ent: WordNodeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wngb *WordNodeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wngb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wngb *WordNodeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wngb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wordnode.Label}
	default:
		err = fmt.Errorf("ent: WordNodeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wngb *WordNodeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wngb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (wngb *WordNodeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wngb.fields) > 1 {
		return nil, errors.New("ent: WordNodeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wngb *WordNodeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wngb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (wngb *WordNodeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wngb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wordnode.Label}
	default:
		err = fmt.Errorf("ent: WordNodeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wngb *WordNodeGroupBy) BoolX(ctx context.Context) bool {
	v, err := wngb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wngb *WordNodeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range wngb.fields {
		if !wordnode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := wngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wngb *WordNodeGroupBy) sqlQuery() *sql.Selector {
	selector := wngb.sql.Select()
	aggregation := make([]string, 0, len(wngb.fns))
	for _, fn := range wngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(wngb.fields)+len(wngb.fns))
		for _, f := range wngb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(wngb.fields...)...)
}

// WordNodeSelect is the builder for selecting fields of WordNode entities.
type WordNodeSelect struct {
	*WordNodeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (wns *WordNodeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := wns.prepareQuery(ctx); err != nil {
		return err
	}
	wns.sql = wns.WordNodeQuery.sqlQuery(ctx)
	return wns.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wns *WordNodeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wns.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (wns *WordNodeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wns.fields) > 1 {
		return nil, errors.New("ent: WordNodeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wns *WordNodeSelect) StringsX(ctx context.Context) []string {
	v, err := wns.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (wns *WordNodeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wns.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wordnode.Label}
	default:
		err = fmt.Errorf("ent: WordNodeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wns *WordNodeSelect) StringX(ctx context.Context) string {
	v, err := wns.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (wns *WordNodeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wns.fields) > 1 {
		return nil, errors.New("ent: WordNodeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wns *WordNodeSelect) IntsX(ctx context.Context) []int {
	v, err := wns.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (wns *WordNodeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wns.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wordnode.Label}
	default:
		err = fmt.Errorf("ent: WordNodeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wns *WordNodeSelect) IntX(ctx context.Context) int {
	v, err := wns.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (wns *WordNodeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wns.fields) > 1 {
		return nil, errors.New("ent: WordNodeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wns *WordNodeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wns.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (wns *WordNodeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wns.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wordnode.Label}
	default:
		err = fmt.Errorf("ent: WordNodeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wns *WordNodeSelect) Float64X(ctx context.Context) float64 {
	v, err := wns.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (wns *WordNodeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wns.fields) > 1 {
		return nil, errors.New("ent: WordNodeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wns *WordNodeSelect) BoolsX(ctx context.Context) []bool {
	v, err := wns.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (wns *WordNodeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wns.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{wordnode.Label}
	default:
		err = fmt.Errorf("ent: WordNodeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wns *WordNodeSelect) BoolX(ctx context.Context) bool {
	v, err := wns.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wns *WordNodeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wns.sql.Query()
	if err := wns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
