package testly

import (
	"regexp"
	"strings"
)

func ReplacePlaceholders(template string, variables map[string]interface{}) string {
	re := regexp.MustCompile(`\$\{\{\s*env\.(\w+)\s*\}\}`) // Regex to match `${{ env.variable }}`

	return re.ReplaceAllStringFunc(template, func(match string) string {
		// Extract variable name (without `env.`)
		key := strings.TrimSpace(match[7 : len(match)-2]) // Remove `${{ env. }}`

		// Replace with value if it exists, otherwise keep original placeholder
		if val, exists := variables[key]; exists {
			return val.(string)
		}
		return match // Return original placeholder if not found
	})
}
