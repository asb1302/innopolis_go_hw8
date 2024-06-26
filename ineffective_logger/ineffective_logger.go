package ineffective_logger

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type IneffectiveLogger struct {
	mu sync.Mutex
}

func (l *IneffectiveLogger) Info(msg string) {
	// Несколько этапов форматирования строки
	timestamp := time.Now().Format(time.RFC3339)
	formattedMsg := fmt.Sprintf("%s: %s\n", timestamp, msg)

	// Излишнее копирование данных
	var finalMsg string
	for _, char := range formattedMsg {
		finalMsg += string(char)
	}

	// Использование неэффективных структур данных
	data := []byte(finalMsg)
	builder := strings.Builder{}
	for _, b := range data {
		builder.WriteByte(b)
	}

	l.mu.Lock()
	fmt.Print(finalMsg)
	l.mu.Unlock()
}
