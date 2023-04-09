package logger

import (
	"fmt"
	"time"
)

type StdoutLogger struct {
	Loggerable
}

func NewStdoutLogger(name string) *StdoutLogger {
	return &StdoutLogger{
		Loggerable: func(level Level, message string) error {
			fmt.Println(
				fmt.Sprintf(
					"[%s] %s.%s: %s",
					time.Now().Format("2006-01-02 15:04:05"),
					name,
					level.UpperString(), message,
				),
			)

			return nil
		},
	}
}
