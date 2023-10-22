package http_api

import (
	"log"
	"runtime"
)

// Log errors
func error_logging(err error) {
	if err != nil {
		pc := make([]uintptr, 10)
		n := runtime.Callers(2, pc)
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		// fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		log.Printf("[API] error on %s: %s", frame.Function, err)
	}
}
