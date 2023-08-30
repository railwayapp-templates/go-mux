package tools

import "os"

// Returns PORT from environment if found, defaults to
// value in `port` parameter otherwise. The returned port
// is prefixed with a `:`, e.g. `":3000"`.
func EnvPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
