package kctx

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type EchoContext struct {
	echo.Context
}

func (c *EchoContext) GetLog() *zap.Logger {
	return c.Get("log").(*zap.Logger)
}

func (c *EchoContext) Ctx() context.Context {
	return context.Background()
}

func (c *EchoContext) Extend() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &EchoContext{c}
			return next(cc)
		}
	}
}
