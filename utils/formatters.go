package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// GetBranchName
// get a correctly formattet branch
func GetBranchName(raw string, issue string) string {

	raw = strings.Join(strings.Fields(strings.ToLower(raw)), "-")
	m := regexp.MustCompile("![a-z0-9]")
	mn := regexp.MustCompile(`\.`)
	cleanedName := m.ReplaceAllString(raw, "")
	cleanedName = mn.ReplaceAllString(cleanedName, "-")

	return fmt.Sprintf("%s-%s", issue, cleanedName)
}
