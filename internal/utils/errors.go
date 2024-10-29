package utils

import (
	"fmt"
	"log"
)

type CustomError struct {
	Module   string // Module where the error occurred (e.g., "Scraper", "Storage")
	Function string // Function where the error originated
	Err      error  // The original error
	Severity string // Severity level ("INFO", "WARNING", "ERROR")
	Message  string // Custom message for additional context
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%s] %s - %s: %v", e.Severity, e.Module, e.Message, e.Err)
}

func NewCustomError(module, function, severity, message string, err error) *CustomError {
	return &CustomError{
		Module:   module,
		Function: function,
		Err:      err,
		Severity: severity,
		Message:  message,
	}
}

func LogError(e *CustomError) {
	if e.Severity == "ERROR" || e.Severity == "WARNING" {
		log.Printf("ERROR: %s", e.Error())
	} else {
		log.Printf("INFO: %s", e.Error())
	}
}
