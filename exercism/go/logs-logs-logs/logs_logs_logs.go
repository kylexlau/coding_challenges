package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		if char == '❗' {
			return "recommendation"
		} else if char == '🔍' {
			return "search"
		} else if char == '☀' {
			return "weather"
		}
	}
	return "default"

}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	log_rune := []rune(log)
	for i, c := range log_rune {
		if c == oldRune {
			log_rune[i] = newRune
		}
	}
	return string(log_rune)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	length := utf8.RuneCountInString(log)
	if length <= limit {
		return true
	}
	return false
}
