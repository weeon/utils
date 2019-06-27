package ctxutil

import (
	"context"

	"github.com/weeon/contract"
	"github.com/weeon/utils"
)

func GetRequestIDFromContext(c context.Context) string {
	v, ok := c.Value(contract.RequestID).(string)
	if !ok {
		return "RequestID not found"
	}
	return v
}

func AddRequestID(c context.Context, reqID ...string) context.Context {
	if reqID != nil && len(reqID) != 0 {
		return context.WithValue(c, contract.RequestID, reqID[0])
	}
	return context.WithValue(c, contract.RequestID, utils.NewUUID())
}

func CopyRequestID(ctx context.Context, fromCtx context.Context) context.Context {
	return context.WithValue(ctx, contract.RequestID, GetRequestIDFromContext(fromCtx))
}
