package middleware

import (
	"github.com/Unknwon/macaron"
)

// Context represents context of a request.
type Context struct {
	*macaron.Context
}

func Contexter() macaron.Handler {
	return func(c *macaron.Context) {
		ctx := &Context{
			Context: c,
		}
		c.Map(ctx)
	}
}
