package modules

import (
	"regexp"
	"strings"
)

// ValidateModuleName returns true if it's a valid module name.
// It allows any number of / and . characters in between.
func ValidateModuleName(moduleName string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+(?:[\\/.][a-zA-Z0-9_-]+)*$", moduleName)
	return matched
}

// GetRootDir returns the project directory name from the module path.
// Returns the last token by splitting the moduleName with /
func GetRootDir(moduleName string) string {
	tokens := strings.Split(moduleName, "/")
	return tokens[len(tokens)-1]
}
