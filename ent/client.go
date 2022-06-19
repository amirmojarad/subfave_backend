// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"interface_project/ent/migrate"

	"interface_project/ent/collection"
	"interface_project/ent/fileentity"
	"interface_project/ent/movie"
	"interface_project/ent/searchkeyword"
	"interface_project/ent/user"
	"interface_project/ent/word"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Collection is the client for interacting with the Collection builders.
	Collection *CollectionClient
	// FileEntity is the client for interacting with the FileEntity builders.
	FileEntity *FileEntityClient
	// Movie is the client for interacting with the Movie builders.
	Movie *MovieClient
	// SearchKeyword is the client for interacting with the SearchKeyword builders.
	SearchKeyword *SearchKeywordClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Word is the client for interacting with the Word builders.
	Word *WordClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Collection = NewCollectionClient(c.config)
	c.FileEntity = NewFileEntityClient(c.config)
	c.Movie = NewMovieClient(c.config)
	c.SearchKeyword = NewSearchKeywordClient(c.config)
	c.User = NewUserClient(c.config)
	c.Word = NewWordClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Collection:    NewCollectionClient(cfg),
		FileEntity:    NewFileEntityClient(cfg),
		Movie:         NewMovieClient(cfg),
		SearchKeyword: NewSearchKeywordClient(cfg),
		User:          NewUserClient(cfg),
		Word:          NewWordClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Collection:    NewCollectionClient(cfg),
		FileEntity:    NewFileEntityClient(cfg),
		Movie:         NewMovieClient(cfg),
		SearchKeyword: NewSearchKeywordClient(cfg),
		User:          NewUserClient(cfg),
		Word:          NewWordClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Collection.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Collection.Use(hooks...)
	c.FileEntity.Use(hooks...)
	c.Movie.Use(hooks...)
	c.SearchKeyword.Use(hooks...)
	c.User.Use(hooks...)
	c.Word.Use(hooks...)
}

// CollectionClient is a client for the Collection schema.
type CollectionClient struct {
	config
}

// NewCollectionClient returns a client for the Collection from the given config.
func NewCollectionClient(c config) *CollectionClient {
	return &CollectionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `collection.Hooks(f(g(h())))`.
func (c *CollectionClient) Use(hooks ...Hook) {
	c.hooks.Collection = append(c.hooks.Collection, hooks...)
}

// Create returns a create builder for Collection.
func (c *CollectionClient) Create() *CollectionCreate {
	mutation := newCollectionMutation(c.config, OpCreate)
	return &CollectionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Collection entities.
func (c *CollectionClient) CreateBulk(builders ...*CollectionCreate) *CollectionCreateBulk {
	return &CollectionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Collection.
func (c *CollectionClient) Update() *CollectionUpdate {
	mutation := newCollectionMutation(c.config, OpUpdate)
	return &CollectionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CollectionClient) UpdateOne(co *Collection) *CollectionUpdateOne {
	mutation := newCollectionMutation(c.config, OpUpdateOne, withCollection(co))
	return &CollectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CollectionClient) UpdateOneID(id int) *CollectionUpdateOne {
	mutation := newCollectionMutation(c.config, OpUpdateOne, withCollectionID(id))
	return &CollectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Collection.
func (c *CollectionClient) Delete() *CollectionDelete {
	mutation := newCollectionMutation(c.config, OpDelete)
	return &CollectionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CollectionClient) DeleteOne(co *Collection) *CollectionDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CollectionClient) DeleteOneID(id int) *CollectionDeleteOne {
	builder := c.Delete().Where(collection.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CollectionDeleteOne{builder}
}

// Query returns a query builder for Collection.
func (c *CollectionClient) Query() *CollectionQuery {
	return &CollectionQuery{
		config: c.config,
	}
}

// Get returns a Collection entity by its id.
func (c *CollectionClient) Get(ctx context.Context, id int) (*Collection, error) {
	return c.Query().Where(collection.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CollectionClient) GetX(ctx context.Context, id int) *Collection {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Collection.
func (c *CollectionClient) QueryUser(co *Collection) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(collection.Table, collection.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, collection.UserTable, collection.UserPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCollectionWords queries the collection_words edge of a Collection.
func (c *CollectionClient) QueryCollectionWords(co *Collection) *WordQuery {
	query := &WordQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(collection.Table, collection.FieldID, id),
			sqlgraph.To(word.Table, word.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, collection.CollectionWordsTable, collection.CollectionWordsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CollectionClient) Hooks() []Hook {
	return c.hooks.Collection
}

// FileEntityClient is a client for the FileEntity schema.
type FileEntityClient struct {
	config
}

// NewFileEntityClient returns a client for the FileEntity from the given config.
func NewFileEntityClient(c config) *FileEntityClient {
	return &FileEntityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `fileentity.Hooks(f(g(h())))`.
func (c *FileEntityClient) Use(hooks ...Hook) {
	c.hooks.FileEntity = append(c.hooks.FileEntity, hooks...)
}

// Create returns a create builder for FileEntity.
func (c *FileEntityClient) Create() *FileEntityCreate {
	mutation := newFileEntityMutation(c.config, OpCreate)
	return &FileEntityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FileEntity entities.
func (c *FileEntityClient) CreateBulk(builders ...*FileEntityCreate) *FileEntityCreateBulk {
	return &FileEntityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FileEntity.
func (c *FileEntityClient) Update() *FileEntityUpdate {
	mutation := newFileEntityMutation(c.config, OpUpdate)
	return &FileEntityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FileEntityClient) UpdateOne(fe *FileEntity) *FileEntityUpdateOne {
	mutation := newFileEntityMutation(c.config, OpUpdateOne, withFileEntity(fe))
	return &FileEntityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FileEntityClient) UpdateOneID(id int) *FileEntityUpdateOne {
	mutation := newFileEntityMutation(c.config, OpUpdateOne, withFileEntityID(id))
	return &FileEntityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FileEntity.
func (c *FileEntityClient) Delete() *FileEntityDelete {
	mutation := newFileEntityMutation(c.config, OpDelete)
	return &FileEntityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FileEntityClient) DeleteOne(fe *FileEntity) *FileEntityDeleteOne {
	return c.DeleteOneID(fe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FileEntityClient) DeleteOneID(id int) *FileEntityDeleteOne {
	builder := c.Delete().Where(fileentity.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FileEntityDeleteOne{builder}
}

// Query returns a query builder for FileEntity.
func (c *FileEntityClient) Query() *FileEntityQuery {
	return &FileEntityQuery{
		config: c.config,
	}
}

// Get returns a FileEntity entity by its id.
func (c *FileEntityClient) Get(ctx context.Context, id int) (*FileEntity, error) {
	return c.Query().Where(fileentity.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FileEntityClient) GetX(ctx context.Context, id int) *FileEntity {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a FileEntity.
func (c *FileEntityClient) QueryOwner(fe *FileEntity) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(fileentity.Table, fileentity.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, fileentity.OwnerTable, fileentity.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(fe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWords queries the words edge of a FileEntity.
func (c *FileEntityClient) QueryWords(fe *FileEntity) *WordQuery {
	query := &WordQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := fe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(fileentity.Table, fileentity.FieldID, id),
			sqlgraph.To(word.Table, word.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, fileentity.WordsTable, fileentity.WordsColumn),
		)
		fromV = sqlgraph.Neighbors(fe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FileEntityClient) Hooks() []Hook {
	return c.hooks.FileEntity
}

// MovieClient is a client for the Movie schema.
type MovieClient struct {
	config
}

// NewMovieClient returns a client for the Movie from the given config.
func NewMovieClient(c config) *MovieClient {
	return &MovieClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `movie.Hooks(f(g(h())))`.
func (c *MovieClient) Use(hooks ...Hook) {
	c.hooks.Movie = append(c.hooks.Movie, hooks...)
}

// Create returns a create builder for Movie.
func (c *MovieClient) Create() *MovieCreate {
	mutation := newMovieMutation(c.config, OpCreate)
	return &MovieCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Movie entities.
func (c *MovieClient) CreateBulk(builders ...*MovieCreate) *MovieCreateBulk {
	return &MovieCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Movie.
func (c *MovieClient) Update() *MovieUpdate {
	mutation := newMovieMutation(c.config, OpUpdate)
	return &MovieUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MovieClient) UpdateOne(m *Movie) *MovieUpdateOne {
	mutation := newMovieMutation(c.config, OpUpdateOne, withMovie(m))
	return &MovieUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MovieClient) UpdateOneID(id int) *MovieUpdateOne {
	mutation := newMovieMutation(c.config, OpUpdateOne, withMovieID(id))
	return &MovieUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Movie.
func (c *MovieClient) Delete() *MovieDelete {
	mutation := newMovieMutation(c.config, OpDelete)
	return &MovieDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MovieClient) DeleteOne(m *Movie) *MovieDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MovieClient) DeleteOneID(id int) *MovieDeleteOne {
	builder := c.Delete().Where(movie.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MovieDeleteOne{builder}
}

// Query returns a query builder for Movie.
func (c *MovieClient) Query() *MovieQuery {
	return &MovieQuery{
		config: c.config,
	}
}

// Get returns a Movie entity by its id.
func (c *MovieClient) Get(ctx context.Context, id int) (*Movie, error) {
	return c.Query().Where(movie.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MovieClient) GetX(ctx context.Context, id int) *Movie {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Movie.
func (c *MovieClient) QueryUsers(m *Movie) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(movie.Table, movie.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, movie.UsersTable, movie.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MovieClient) Hooks() []Hook {
	return c.hooks.Movie
}

// SearchKeywordClient is a client for the SearchKeyword schema.
type SearchKeywordClient struct {
	config
}

// NewSearchKeywordClient returns a client for the SearchKeyword from the given config.
func NewSearchKeywordClient(c config) *SearchKeywordClient {
	return &SearchKeywordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `searchkeyword.Hooks(f(g(h())))`.
func (c *SearchKeywordClient) Use(hooks ...Hook) {
	c.hooks.SearchKeyword = append(c.hooks.SearchKeyword, hooks...)
}

// Create returns a create builder for SearchKeyword.
func (c *SearchKeywordClient) Create() *SearchKeywordCreate {
	mutation := newSearchKeywordMutation(c.config, OpCreate)
	return &SearchKeywordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SearchKeyword entities.
func (c *SearchKeywordClient) CreateBulk(builders ...*SearchKeywordCreate) *SearchKeywordCreateBulk {
	return &SearchKeywordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SearchKeyword.
func (c *SearchKeywordClient) Update() *SearchKeywordUpdate {
	mutation := newSearchKeywordMutation(c.config, OpUpdate)
	return &SearchKeywordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SearchKeywordClient) UpdateOne(sk *SearchKeyword) *SearchKeywordUpdateOne {
	mutation := newSearchKeywordMutation(c.config, OpUpdateOne, withSearchKeyword(sk))
	return &SearchKeywordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SearchKeywordClient) UpdateOneID(id int) *SearchKeywordUpdateOne {
	mutation := newSearchKeywordMutation(c.config, OpUpdateOne, withSearchKeywordID(id))
	return &SearchKeywordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SearchKeyword.
func (c *SearchKeywordClient) Delete() *SearchKeywordDelete {
	mutation := newSearchKeywordMutation(c.config, OpDelete)
	return &SearchKeywordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SearchKeywordClient) DeleteOne(sk *SearchKeyword) *SearchKeywordDeleteOne {
	return c.DeleteOneID(sk.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SearchKeywordClient) DeleteOneID(id int) *SearchKeywordDeleteOne {
	builder := c.Delete().Where(searchkeyword.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SearchKeywordDeleteOne{builder}
}

// Query returns a query builder for SearchKeyword.
func (c *SearchKeywordClient) Query() *SearchKeywordQuery {
	return &SearchKeywordQuery{
		config: c.config,
	}
}

// Get returns a SearchKeyword entity by its id.
func (c *SearchKeywordClient) Get(ctx context.Context, id int) (*SearchKeyword, error) {
	return c.Query().Where(searchkeyword.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SearchKeywordClient) GetX(ctx context.Context, id int) *SearchKeyword {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a SearchKeyword.
func (c *SearchKeywordClient) QueryUser(sk *SearchKeyword) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := sk.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(searchkeyword.Table, searchkeyword.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, searchkeyword.UserTable, searchkeyword.UserColumn),
		)
		fromV = sqlgraph.Neighbors(sk.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SearchKeywordClient) Hooks() []Hook {
	return c.hooks.SearchKeyword
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFavoriteMovies queries the favorite_movies edge of a User.
func (c *UserClient) QueryFavoriteMovies(u *User) *MovieQuery {
	query := &MovieQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.FavoriteMoviesTable, user.FavoriteMoviesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySearchedKeywords queries the searched_keywords edge of a User.
func (c *UserClient) QuerySearchedKeywords(u *User) *SearchKeywordQuery {
	query := &SearchKeywordQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(searchkeyword.Table, searchkeyword.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.SearchedKeywordsTable, user.SearchedKeywordsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFavoriteWords queries the favorite_words edge of a User.
func (c *UserClient) QueryFavoriteWords(u *User) *WordQuery {
	query := &WordQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(word.Table, word.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.FavoriteWordsTable, user.FavoriteWordsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFiles queries the files edge of a User.
func (c *UserClient) QueryFiles(u *User) *FileEntityQuery {
	query := &FileEntityQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(fileentity.Table, fileentity.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.FilesTable, user.FilesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCollections queries the collections edge of a User.
func (c *UserClient) QueryCollections(u *User) *CollectionQuery {
	query := &CollectionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(collection.Table, collection.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.CollectionsTable, user.CollectionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWords queries the words edge of a User.
func (c *UserClient) QueryWords(u *User) *WordQuery {
	query := &WordQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(word.Table, word.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.WordsTable, user.WordsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// WordClient is a client for the Word schema.
type WordClient struct {
	config
}

// NewWordClient returns a client for the Word from the given config.
func NewWordClient(c config) *WordClient {
	return &WordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `word.Hooks(f(g(h())))`.
func (c *WordClient) Use(hooks ...Hook) {
	c.hooks.Word = append(c.hooks.Word, hooks...)
}

// Create returns a create builder for Word.
func (c *WordClient) Create() *WordCreate {
	mutation := newWordMutation(c.config, OpCreate)
	return &WordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Word entities.
func (c *WordClient) CreateBulk(builders ...*WordCreate) *WordCreateBulk {
	return &WordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Word.
func (c *WordClient) Update() *WordUpdate {
	mutation := newWordMutation(c.config, OpUpdate)
	return &WordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WordClient) UpdateOne(w *Word) *WordUpdateOne {
	mutation := newWordMutation(c.config, OpUpdateOne, withWord(w))
	return &WordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WordClient) UpdateOneID(id int) *WordUpdateOne {
	mutation := newWordMutation(c.config, OpUpdateOne, withWordID(id))
	return &WordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Word.
func (c *WordClient) Delete() *WordDelete {
	mutation := newWordMutation(c.config, OpDelete)
	return &WordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WordClient) DeleteOne(w *Word) *WordDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WordClient) DeleteOneID(id int) *WordDeleteOne {
	builder := c.Delete().Where(word.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WordDeleteOne{builder}
}

// Query returns a query builder for Word.
func (c *WordClient) Query() *WordQuery {
	return &WordQuery{
		config: c.config,
	}
}

// Get returns a Word entity by its id.
func (c *WordClient) Get(ctx context.Context, id int) (*Word, error) {
	return c.Query().Where(word.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WordClient) GetX(ctx context.Context, id int) *Word {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Word.
func (c *WordClient) QueryUser(w *Word) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, word.UserTable, word.UserColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFile queries the file edge of a Word.
func (c *WordClient) QueryFile(w *Word) *FileEntityQuery {
	query := &FileEntityQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, id),
			sqlgraph.To(fileentity.Table, fileentity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, word.FileTable, word.FileColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCollection queries the collection edge of a Word.
func (c *WordClient) QueryCollection(w *Word) *CollectionQuery {
	query := &CollectionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, id),
			sqlgraph.To(collection.Table, collection.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, word.CollectionTable, word.CollectionPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwner queries the owner edge of a Word.
func (c *WordClient) QueryOwner(w *Word) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, word.OwnerTable, word.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WordClient) Hooks() []Hook {
	return c.hooks.Word
}
