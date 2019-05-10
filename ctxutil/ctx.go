package ctxutil

import (
	"context"
	"github.com/weeon/contract"
)

func GetRequestIDFromContext(c context.Context) string {
	v, ok := c.Value(contract.RequestID).(string)
	if !ok {
		return "RequestID not found"
	}
	return v
}
