package parsinglogfiles

import (
	"regexp"
)

var LogSeparatorRe = regexp.MustCompile("<[*-=~]*>")

var LogLineRe = regexp.MustCompile(`^\[(WRN|TRC|DBG|INF|ERR|FTL)\]`)

var LogEOLRe = regexp.MustCompile(`(?i)end-of-line\w*`)

var LogPasswordRe = regexp.MustCompile("(?i)\".*password.*\"")

var LogUserRe = regexp.MustCompile(`User\s+(\w+)`)

func IsValidLine(text string) bool {
	return LogLineRe.MatchString(text)
}

func SplitLogLine(text string) []string {
	return LogSeparatorRe.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var counter int
	for _, line := range lines {
		if LogPasswordRe.Match([]byte(line)) {
			counter++
		}
	}

	return counter
}

func RemoveEndOfLineText(text string) string {
	result := LogEOLRe.ReplaceAll([]byte(text), []byte{})
	return string(result)
}

func TagWithUserName(lines []string) []string {
	result := make([]string, len(lines))
	copy(result, lines)

	for i, line := range result {
		res := LogUserRe.FindStringSubmatch(line)

		if res != nil {
			result[i] = "[USR] " + res[1] + " " + result[i]
		}
	}

	return result
}
