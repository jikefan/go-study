package context

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

type key struct{}

var k = key{}

func NewRequestID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func NewContextWithTraceID() context.Context {
	return context.WithValue(context.Background(), k, NewRequestID())
}

func PrintLog(ctx context.Context, message string) {
	fmt.Printf("%s|info|trace_id=%s|%s", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, k), message)
}

func GetContextValue(ctx context.Context, k key) string {
	v, ok := ctx.Value(k).(string)
	if !ok {
		return ""
	}

	return v
}

func TestWithValue(t *testing.T) {
	ctx := NewContextWithTraceID()
	PrintLog(ctx, "Go context WithValue")
}


