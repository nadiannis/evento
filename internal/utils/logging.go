package utils

import "context"

type contextKey string

const LogMessageCtxKey contextKey = contextKey("logMessage")

func SetLogMessage(ctx context.Context, message any) context.Context {
	return context.WithValue(ctx, LogMessageCtxKey, message)
}

func GetLogMessage(ctx context.Context) any {
	return ctx.Value(LogMessageCtxKey)
}
