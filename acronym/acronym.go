package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate takes a string and returns its acroynm
func Abbreviate(s string) string {
	re := regexp.MustCompile(`[^\sa-zA-Z0-9\-]+`)
	s = re.ReplaceAllString(s, "")
	re = regexp.MustCompile(`\b[A-Z]`)
	letters := re.FindAllString(strings.Title(strings.ToLower(s)), -1)
	return strings.Join(letters[:], "")
}
