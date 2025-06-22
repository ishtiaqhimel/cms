package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ToP[T comparable](p T) *T {
	return &p
}

var (
	// Compile regexes once to avoid recompilation on every call
	reInvalidChars = regexp.MustCompile(`[^\p{L}\p{M}\p{N}\s-]+`)
	reTrimSpaces   = regexp.MustCompile(`^\s+|\s+$`)
	reHyphenate    = regexp.MustCompile(`\s*-+\s*|\s+`)
)

func GenerateSlug(s string) (string, error) {
	if len(strings.TrimSpace(s)) == 0 {
		return "", fmt.Errorf("can't slugify empty string")
	}

	// Replace `&` with `and`
	s = strings.ReplaceAll(s, "&", " and ")

	// Remove invalid characters (non-letter/number/space/hyphen)
	s = reInvalidChars.ReplaceAllString(s, "")
	// Trim leading/trailing spaces
	s = reTrimSpaces.ReplaceAllString(s, "")
	// Replace spaces and hyphens with single hyphen
	s = reHyphenate.ReplaceAllString(s, "-")

	return strings.ToLower(s), nil
}
