package product

import "strings"

// GenerateSlug creates a URL-friendly slug from a product name
func GenerateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
