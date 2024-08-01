package effective_logger

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type EffectiveLogger struct {
	mu sync.Mutex
}

func (l *EffectiveLogger) Info(msg string) {
	var builder strings.Builder

	builder.WriteString(time.Now().Format(time.RFC3339))
	builder.WriteString(": ")
	builder.WriteString(msg)
	builder.WriteString("\n")

	l.mu.Lock()
	fmt.Print(builder.String())
	l.mu.Unlock()
}
