package util

import "context"

type ctxKeys uint8

const (
	// UsernameCtxKey is key that can be used to retrive username from context
	UsernameCtxKey ctxKeys = iota + 1
)

// WrapCtx : wraps context with key value
func WrapCtx(ctx context.Context, key ctxKeys, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

// UnWrapCtx : un-wraps context can return value of the key
func UnWrapCtx(ctx context.Context, key ctxKeys) interface{} {
	return ctx.Value(key)
}
