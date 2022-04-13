package config

import (
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/meta"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

const (
	DefaultInterval = time.Minute
)

type config struct {
	endpoint string
	database string
	secure   bool
	meta     meta.Meta

	operationTimeout     time.Duration
	operationCancelAfter time.Duration

	interval time.Duration
	trace    trace.Discovery

	panicCallback func(e interface{})
}

func (c *config) PanicCallback() func(e interface{}) {
	return c.panicCallback
}

func (c *config) Meta() meta.Meta {
	return c.meta
}

func (c *config) OperationTimeout() time.Duration {
	return c.operationTimeout
}

func (c *config) OperationCancelAfter() time.Duration {
	return c.operationCancelAfter
}

func (c *config) Interval() time.Duration {
	return c.interval
}

func (c *config) Endpoint() string {
	return c.endpoint
}

func (c *config) Database() string {
	return c.database
}

func (c *config) Secure() bool {
	return c.secure
}

func (c *config) Trace() trace.Discovery {
	return c.trace
}

type Option func(c *config)

// WithEndpoint set a required starting endpoint for connect
func WithEndpoint(endpoint string) Option {
	return func(c *config) {
		c.endpoint = endpoint
	}
}

// WithDatabase set a required database name.
func WithDatabase(database string) Option {
	return func(c *config) {
		c.database = database
	}
}

// WithSecure set flag for secure connection
func WithSecure(ssl bool) Option {
	return func(c *config) {
		c.secure = ssl
	}
}

// WithMeta is not for user.
//
// This option add meta information about database connection
func WithMeta(meta meta.Meta) Option {
	return func(c *config) {
		c.meta = meta
	}
}

// WithTrace configures discovery client calls tracing
func WithTrace(trace trace.Discovery, opts ...trace.DiscoveryComposeOption) Option {
	return func(c *config) {
		c.trace = c.trace.Compose(trace, opts...)
	}
}

// WithOperationTimeout define the maximum amount of time a YDB server will process
// an operation. After timeout exceeds YDB will try to cancel operation and
// regardless of the cancellation appropriate error will be returned to
// the client.
//
// If OperationTimeout is zero then no timeout is used.
func WithOperationTimeout(operationTimeout time.Duration) Option {
	return func(c *config) {
		c.operationTimeout = operationTimeout
	}
}

// WithOperationCancelAfter set the maximum amount of time a YDB server will process an
// operation. After timeout exceeds YDB will try to cancel operation and if
// it succeeds appropriate error will be returned to the client; otherwise
// processing will be continued.
//
// If OperationCancelAfter is zero then no timeout is used.
func WithOperationCancelAfter(operationCancelAfter time.Duration) Option {
	return func(c *config) {
		c.operationCancelAfter = operationCancelAfter
	}
}

// WithInterval set the frequency of background tasks of ydb endpoints discovery.
//
// If Interval is zero then the DefaultInterval is used.
//
// If Interval is negative, then no background discovery prepared.
func WithInterval(interval time.Duration) Option {
	return func(c *config) {
		if interval <= 0 {
			c.interval = 0
		} else {
			c.interval = interval
		}
	}
}

// WithPanicCallback adds user-defined panic callback
//
// nil will turn off callback
func WithPanicCallback(cb func(e interface{})) Option {
	return func(c *config) {
		c.panicCallback = cb
	}
}

func New(opts ...Option) *config {
	c := &config{
		interval: DefaultInterval,
	}
	for _, o := range opts {
		o(c)
	}
	return c
}
