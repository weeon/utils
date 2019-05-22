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

func AddRequestID(c context.Context) context.Context {
	return context.WithValue(c, contract.RequestID, utils.NewUUID())
}
