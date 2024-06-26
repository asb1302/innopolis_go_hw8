package ineffective_logger

import (
	"testing"
)

func BenchmarkInEffectiveLogger(b *testing.B) {
	logger := &IneffectiveLogger{}
	for i := 0; i < b.N; i++ {
		logger.Info("test message")
	}
}
