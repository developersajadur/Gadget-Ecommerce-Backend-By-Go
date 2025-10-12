package helpers

import (
	"regexp"
	"strings"
)

func GenerateSlug(name string) string {
	slug := strings.ToLower(name)

	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}
