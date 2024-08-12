package logs

import "unicode/utf8"

const (
	LogIdentificatorRecommendation = '❗'
	LogIdentificatorSearch         = '🔍'
	LogIdentificatorWeather        = '☀'
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		switch char {
		case LogIdentificatorRecommendation:
			return "recommendation"
		case LogIdentificatorSearch:
			return "search"
		case LogIdentificatorWeather:
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	word := make([]rune, 0, utf8.RuneCountInString(log))

	for _, char := range log {
		if char == oldRune {
			word = append(word, newRune)
		} else {
			word = append(word, char)
		}
	}

	return string(word)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
