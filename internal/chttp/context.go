package chttp

import (
	"context"
	"time"
)

var untilTimeOut int = 300

func getTimedContext(ms int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(ms)*time.Millisecond)
}
