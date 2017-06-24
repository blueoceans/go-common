package errors

import (
	"strings"
)

// ContainsErrorMessage returns is whether it contains messages or not.
func ContainsErrorMessage(
	err error,
	messages []string,
) bool {
	if err == nil {
		return false
	}
	errorMessage := err.Error()
	for _, message := range messages {
		if strings.Contains(errorMessage, message) {
			return true
		}
	}
	return false
}
