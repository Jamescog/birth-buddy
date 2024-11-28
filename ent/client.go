// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Jamescog/birth-buddy/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/Jamescog/birth-buddy/ent/users"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Users is the client for interacting with the Users builders.
	Users *UsersClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Users = NewUsersClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Users:  NewUsersClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:    ctx,
		config: cfg,
		Users:  NewUsersClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Users.
//		Query().
//		Count(ctx)
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
	c.Users.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Users.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *UsersMutation:
		return c.Users.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// UsersClient is a client for the Users schema.
type UsersClient struct {
	config
}

// NewUsersClient returns a client for the Users from the given config.
func NewUsersClient(c config) *UsersClient {
	return &UsersClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `users.Hooks(f(g(h())))`.
func (c *UsersClient) Use(hooks ...Hook) {
	c.hooks.Users = append(c.hooks.Users, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `users.Intercept(f(g(h())))`.
func (c *UsersClient) Intercept(interceptors ...Interceptor) {
	c.inters.Users = append(c.inters.Users, interceptors...)
}

// Create returns a builder for creating a Users entity.
func (c *UsersClient) Create() *UsersCreate {
	mutation := newUsersMutation(c.config, OpCreate)
	return &UsersCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Users entities.
func (c *UsersClient) CreateBulk(builders ...*UsersCreate) *UsersCreateBulk {
	return &UsersCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UsersClient) MapCreateBulk(slice any, setFunc func(*UsersCreate, int)) *UsersCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UsersCreateBulk{err: fmt.Errorf("calling to UsersClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UsersCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UsersCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Users.
func (c *UsersClient) Update() *UsersUpdate {
	mutation := newUsersMutation(c.config, OpUpdate)
	return &UsersUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UsersClient) UpdateOne(u *Users) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsers(u))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UsersClient) UpdateOneID(id int) *UsersUpdateOne {
	mutation := newUsersMutation(c.config, OpUpdateOne, withUsersID(id))
	return &UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Users.
func (c *UsersClient) Delete() *UsersDelete {
	mutation := newUsersMutation(c.config, OpDelete)
	return &UsersDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UsersClient) DeleteOne(u *Users) *UsersDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UsersClient) DeleteOneID(id int) *UsersDeleteOne {
	builder := c.Delete().Where(users.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UsersDeleteOne{builder}
}

// Query returns a query builder for Users.
func (c *UsersClient) Query() *UsersQuery {
	return &UsersQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUsers},
		inters: c.Interceptors(),
	}
}

// Get returns a Users entity by its id.
func (c *UsersClient) Get(ctx context.Context, id int) (*Users, error) {
	return c.Query().Where(users.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UsersClient) GetX(ctx context.Context, id int) *Users {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UsersClient) Hooks() []Hook {
	return c.hooks.Users
}

// Interceptors returns the client interceptors.
func (c *UsersClient) Interceptors() []Interceptor {
	return c.inters.Users
}

func (c *UsersClient) mutate(ctx context.Context, m *UsersMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UsersCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UsersUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UsersUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UsersDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Users mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Users []ent.Hook
	}
	inters struct {
		Users []ent.Interceptor
	}
)
