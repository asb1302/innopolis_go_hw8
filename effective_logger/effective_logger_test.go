package effective_logger

import (
	"testing"
)

func BenchmarkEffectiveLogger(b *testing.B) {
	logger := &EffectiveLogger{}
	for i := 0; i < b.N; i++ {
		logger.Info("test message")
	}
}
