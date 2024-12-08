package utils

import (
	"regexp"
	"strings"
)

func GenerateSlug(name string) string {
	// Convert the string to lowercase
	slug := strings.ToLower(name)

	// Replace spaces and underscores with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")

	// Remove all non-alphanumeric characters except hyphens
	re := regexp.MustCompile("[^a-z0-9-]+")
	slug = re.ReplaceAllString(slug, "")

	// Remove leading and trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}
