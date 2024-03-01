// Code generated by gtrace. DO NOT EDIT.

package trace

import (
	"context"
)

// queryComposeOptions is a holder of options
type queryComposeOptions struct {
	panicCallback func(e interface{})
}

// QueryOption specified Query compose option
type QueryComposeOption func(o *queryComposeOptions)

// WithQueryPanicCallback specified behavior on panic
func WithQueryPanicCallback(cb func(e interface{})) QueryComposeOption {
	return func(o *queryComposeOptions) {
		o.panicCallback = cb
	}
}

// Compose returns a new Query which has functional fields composed both from t and x.
func (t *Query) Compose(x *Query, opts ...QueryComposeOption) *Query {
	var ret Query
	options := queryComposeOptions{}
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	{
		h1 := t.OnDo
		h2 := x.OnDo
		ret.OnDo = func(q QueryDoStartInfo) func(QueryDoIntermediateInfo) func(QueryDoDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(QueryDoIntermediateInfo) func(QueryDoDoneInfo)
			if h1 != nil {
				r = h1(q)
			}
			if h2 != nil {
				r1 = h2(q)
			}
			return func(info QueryDoIntermediateInfo) func(QueryDoDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(QueryDoDoneInfo)
				if r != nil {
					r2 = r(info)
				}
				if r1 != nil {
					r3 = r1(info)
				}
				return func(q QueryDoDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(q)
					}
					if r3 != nil {
						r3(q)
					}
				}
			}
		}
	}
	{
		h1 := t.OnDoTx
		h2 := x.OnDoTx
		ret.OnDoTx = func(q QueryDoTxStartInfo) func(QueryDoTxIntermediateInfo) func(QueryDoTxDoneInfo) {
			if options.panicCallback != nil {
				defer func() {
					if e := recover(); e != nil {
						options.panicCallback(e)
					}
				}()
			}
			var r, r1 func(QueryDoTxIntermediateInfo) func(QueryDoTxDoneInfo)
			if h1 != nil {
				r = h1(q)
			}
			if h2 != nil {
				r1 = h2(q)
			}
			return func(info QueryDoTxIntermediateInfo) func(QueryDoTxDoneInfo) {
				if options.panicCallback != nil {
					defer func() {
						if e := recover(); e != nil {
							options.panicCallback(e)
						}
					}()
				}
				var r2, r3 func(QueryDoTxDoneInfo)
				if r != nil {
					r2 = r(info)
				}
				if r1 != nil {
					r3 = r1(info)
				}
				return func(q QueryDoTxDoneInfo) {
					if options.panicCallback != nil {
						defer func() {
							if e := recover(); e != nil {
								options.panicCallback(e)
							}
						}()
					}
					if r2 != nil {
						r2(q)
					}
					if r3 != nil {
						r3(q)
					}
				}
			}
		}
	}
	return &ret
}
func (t *Query) onDo(q QueryDoStartInfo) func(info QueryDoIntermediateInfo) func(QueryDoDoneInfo) {
	fn := t.OnDo
	if fn == nil {
		return func(QueryDoIntermediateInfo) func(QueryDoDoneInfo) {
			return func(QueryDoDoneInfo) {
				return
			}
		}
	}
	res := fn(q)
	if res == nil {
		return func(QueryDoIntermediateInfo) func(QueryDoDoneInfo) {
			return func(QueryDoDoneInfo) {
				return
			}
		}
	}
	return func(info QueryDoIntermediateInfo) func(QueryDoDoneInfo) {
		res := res(info)
		if res == nil {
			return func(QueryDoDoneInfo) {
				return
			}
		}
		return res
	}
}
func (t *Query) onDoTx(q QueryDoTxStartInfo) func(info QueryDoTxIntermediateInfo) func(QueryDoTxDoneInfo) {
	fn := t.OnDoTx
	if fn == nil {
		return func(QueryDoTxIntermediateInfo) func(QueryDoTxDoneInfo) {
			return func(QueryDoTxDoneInfo) {
				return
			}
		}
	}
	res := fn(q)
	if res == nil {
		return func(QueryDoTxIntermediateInfo) func(QueryDoTxDoneInfo) {
			return func(QueryDoTxDoneInfo) {
				return
			}
		}
	}
	return func(info QueryDoTxIntermediateInfo) func(QueryDoTxDoneInfo) {
		res := res(info)
		if res == nil {
			return func(QueryDoTxDoneInfo) {
				return
			}
		}
		return res
	}
}
func QueryOnDo(t *Query, c *context.Context, call call, label string, idempotent bool, nestedCall bool) func(error) func(attempts int, _ error) {
	var p QueryDoStartInfo
	p.Context = c
	p.Call = call
	p.Label = label
	p.Idempotent = idempotent
	p.NestedCall = nestedCall
	res := t.onDo(p)
	return func(e error) func(int, error) {
		var p QueryDoIntermediateInfo
		p.Error = e
		res := res(p)
		return func(attempts int, e error) {
			var p QueryDoDoneInfo
			p.Attempts = attempts
			p.Error = e
			res(p)
		}
	}
}
func QueryOnDoTx(t *Query, c *context.Context, call call, label string, idempotent bool, nestedCall bool) func(error) func(attempts int, _ error) {
	var p QueryDoTxStartInfo
	p.Context = c
	p.Call = call
	p.Label = label
	p.Idempotent = idempotent
	p.NestedCall = nestedCall
	res := t.onDoTx(p)
	return func(e error) func(int, error) {
		var p QueryDoTxIntermediateInfo
		p.Error = e
		res := res(p)
		return func(attempts int, e error) {
			var p QueryDoTxDoneInfo
			p.Attempts = attempts
			p.Error = e
			res(p)
		}
	}
}
