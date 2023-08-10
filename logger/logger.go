package logger

import (
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"io"
)

type ZapLogger struct {
	*zap.Logger
}

func (z *ZapLogger) Output() io.Writer {
	//TODO implement me
	panic("implement me")
}

func (z *ZapLogger) SetOutput(w io.Writer) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Prefix() string {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) SetPrefix(p string) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Level() log.Lvl {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) SetLevel(v log.Lvl) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) SetHeader(h string) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Print(i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Printf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Printj(j log.JSON) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Debug(i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Debugf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Debugj(j log.JSON) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Info(i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Infof(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Infoj(j log.JSON) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Warn(i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Warnf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Warnj(j log.JSON) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Error(i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Errorf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Errorj(j log.JSON) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Fatal(i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Fatalj(j log.JSON) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Fatalf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Panic(i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Panicj(j log.JSON) {
	//TODO implement me
	panic("implement me")
}

func (z ZapLogger) Panicf(format string, args ...interface{}) {
	//TODO implement me
	panic("implement me")
}
