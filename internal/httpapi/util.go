package httpapi

import (
	"context"

	"github.com/gin-gonic/gin"
)

func wrapUsernameOnCtx(ctx *gin.Context, username string) context.Context {
	return context.WithValue(ctx, "username", username)
}
