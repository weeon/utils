package ginutil

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/weeon/contract"
	"github.com/weeon/utils"
)

func WrapRequestID(c *gin.Context) {
	uuid := utils.NewUUID()
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, contract.RequestID, uuid)
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	SetContext(c, ctx)
	c.Keys[contract.RequestID] = uuid
	c.Header(contract.RequestID, uuid)
	c.Next()
}

func GetRequestID(c *gin.Context) string {
	s, ok := c.Keys[contract.RequestID].(string)
	if !ok {
		return "RequestID Not Found"
	}
	return s
}

func SetContext(c *gin.Context, ctx context.Context) {
	c.Request = c.Request.WithContext(ctx)
}

func GetContext(c *gin.Context) context.Context {
	return c.Request.Context()
}

func GetBearerToken(c *gin.Context, logger contract.Logger) string {
	reqToken := c.GetHeader("Authorization")
	if logger != nil {
		logger.Debugw("Authorization header from request",
			"header", reqToken,
			contract.RequestID, GetRequestID(c),
		)
	}
	splitToken := strings.Split(reqToken, "Bearer")
	if logger != nil {
		logger.Debugw("splitToken",
			"splitTokenArr", splitToken,
			contract.RequestID, GetRequestID(c),
		)
	}
	if len(splitToken) > 1 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
