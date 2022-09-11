package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/micrease/micrease-core/errs"
	"go-micro.dev/v4/server"
)

func RecoverWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
		fmt.Println("RecoverWrapper中件间开始")
		defer errs.Recover(&err)
		err2 := fn(ctx, req, rsp)
		fmt.Println("RecoverWrapper中件间结束", err2)

		if err2 != nil {
			return errors.New(err2.Error())
		}
		return nil
	}
}
