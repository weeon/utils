package ginutil

import (
	"context"

	"git.orx.me/cat/utils"
	"github.com/gin-gonic/gin"
	"github.com/weeon/contract"
)

func WrapRequestID(c *gin.Context) {
	uuid := utils.NewUUID()
	ctx := context.Background()
	ctx = context.WithValue(ctx, contract.RequestID, uuid)
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[contract.Context] = ctx
	c.Keys[contract.RequestID] = uuid
	c.Header(contract.RequestID, uuid)
	c.Next()
}

func GetRequestID(c *gin.Context) string {
	s, ok := c.Keys[contract.RequestID].(string)
	if !ok {
		return "RequestNotFound"
	}
	return s
}

func GetContext(c *gin.Context) context.Context {
	v, ok := c.Keys[contract.Context].(context.Context)
	if !ok {
		return context.Background()
	}
	return v
}
