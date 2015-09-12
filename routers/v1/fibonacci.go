package v1

import (
	"github.com/Unknwon/com"

	"github.com/Unknwon/fws/models"
	"github.com/Unknwon/fws/modules/middleware"
)

func Fibonacci(ctx *middleware.Context) {
	n := ctx.QueryInt("n")
	seq := models.FibonacciSequence(n)
	if seq == nil {
		ctx.JSON(422, map[string]string{
			"error": "Unprocessable number: " + com.ToStr(n),
		})
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"sequence": seq,
	})
}
