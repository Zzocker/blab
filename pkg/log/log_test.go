package log

import (
	"log"
	"testing"
)

type logWriter struct{}

func (lw *logWriter) Write(p []byte) (int, error) {
	return 0, nil
}

func BenchmarkMyLogger(b *testing.B) {
	l := NewLogger(true, &logWriter{}, nil)
	for i := 0; i < b.N; i++ {
		l.Info("%v", nil)
	}
}

func BenchmarkGoLogger(b *testing.B) {
	l := log.New(&logWriter{}, "[INFO]", log.Lshortfile|log.LstdFlags|log.Ltime|log.Ldate|log.Lmicroseconds)
	for i := 0; i < b.N; i++ {
		l.Printf("%v", nil)
	}
}
