package iprange

import (
	"regexp"
	"strings"
)

var (
	ipv4Pattern = regexp.MustCompile(`^(?:25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])(?:\.(?:25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}$`)
	ipv6Pattern = regexp.MustCompile(`^((?:[0-9A-Fa-f]{1,4}:){7}[0-9A-Fa-f]{1,4}|(?:[0-9A-Fa-f]{1,4}:){1,7}:|:(?::[0-9A-Fa-f]{1,4}){1,7}|(?:[0-9A-Fa-f]{1,4}:){1,6}:[0-9A-Fa-f]{1,4}|(?:[0-9A-Fa-f]{1,5}(?::[0-9A-Fa-f]{1,4}){1,2})|(?:[0-9A-Fa-f]{1,4}(?::[0-9A-Fa-f]{1,4}){1,3})|(?:[0-9A-Fa-f]{1,3}(?::[0-9A-Fa-f]{1,4}){1,4})|(?:[0-9A-Fa-f]{1,2}(?::[0-9A-Fa-f]{1,4}){1,5})|[0-9A-Fa-f]{1,4}(?::[0-9A-Fa-f]{1,4}){1,6}|(?:(?:(?:[0-9A-Fa-f]{1,4}:){1,4}|:):(?:[0-9A-Fa-f]{1,4}:){1,4}(?:[0-9A-Fa-f]{1,4}))|::)$`)
)

// IsValid reports whether the provided string is a syntactically valid IPv4 or IPv6 address.
func IsValid(ip string) bool {
	ip = strings.TrimSpace(ip)
	if ip == "" {
		return false
	}

	return ipv4Pattern.MatchString(ip) || ipv6Pattern.MatchString(ip)
}
