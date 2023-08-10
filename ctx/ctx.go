package kctx

import (
	"context"
	"go.uber.org/zap"
)

type Context interface {
	GetLog() *zap.Logger
	Ctx() context.Context
}
