package middleware

import (
	"app/internal/logger"
	"fmt"
	"net"
	"net/http"
	"net/netip"
	"os"
	"strings"

	"golang.org/x/exp/slog"
)

// PrivateRangesCIDR returns a list of private CIDR range
// strings, which can be used as a configuration shortcut.
//
// 192.168.0.0/16, 172.16.0.0/12, 10.0.0.0/8, 127.0.0.1/8, fd00::/8, ::1
func PrivateRanges() []string {
	return []string{
		"192.168.0.0/16",
		"172.16.0.0/12",
		"10.0.0.0/8",
		"127.0.0.1/8",
		"fd00::/8",
		"::1",
	}
}

var proxyIPHeaders = []string{
	"X-Envoy-External-Address",
	"X-Forwarded-For",
	"X-Real-IP",
	"True-Client-IP",
}

var schemeHeaders = []string{
	"X-Forwarded-Proto",
	"X-Forwarded-Scheme",
}

var xForwardedHost = "X-Forwarded-Host"

// TrustProxy checks if the request IP matches one of the provided ranges/IPs
// then inspects common reverse proxy headers and sets the corresponding
// fields in the HTTP request struct for use by middleware or handlers that are next
func TrustProxy(trustedIPs []string) func(http.Handler) http.Handler {
	// parse passed in trusted IPs into a 'netip.Prefix' slice
	parsedIPs, err := parseIPRanges(trustedIPs)
	if err != nil {
		logger.StderrWithSource.Error(err.Error())
		os.Exit(1)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// check if RemoteAddr is trusted
			trusted, err := isTrustedIP(r.RemoteAddr, parsedIPs)
			if err != nil {
				logger.StdoutWithSource.Warn(err.Error(), slog.String("ip", r.RemoteAddr))
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// if RemoteAddr is not trusted, serve next and return early without trusting any proxy headers
			if !trusted {
				next.ServeHTTP(w, r)
				return
			}

			// RemoteAddr is trusted

			// Set the RemoteAddr with the value passed by the proxy
			if realIP := getRealIP(r.Header); realIP != "" {
				r.RemoteAddr = realIP
			}

			// Set the host with the value passed by the proxy
			if host := r.Header.Get(xForwardedHost); host != "" {
				r.Host = host
			}

			// Set the scheme with the value passed by the proxy
			if scheme := getScheme(r.Header); scheme != "" {
				r.URL.Scheme = scheme
			}

			next.ServeHTTP(w, r)
		})
	}
}

// parse passed in IP ranges into a 'netip.Prefix' slice
func parseIPRanges(IPRanges []string) ([]netip.Prefix, error) {
	var parsedIPs []netip.Prefix

	for _, ipStr := range IPRanges {
		if strings.Contains(ipStr, "/") {
			ipNet, err := netip.ParsePrefix(ipStr)
			if err != nil {
				return nil, fmt.Errorf("parsing CIDR expression: %w", err)
			}

			parsedIPs = append(parsedIPs, ipNet)
		} else {
			ipAddr, err := netip.ParseAddr(ipStr)
			if err != nil {
				return nil, fmt.Errorf("invalid IP address: '%s': %w", ipStr, err)
			}

			parsedIPs = append(parsedIPs, netip.PrefixFrom(ipAddr, ipAddr.BitLen()))
		}
	}

	return parsedIPs, nil
}

// check if RemoteAddr is trusted
func isTrustedIP(remoteAddr string, trustedIPs []netip.Prefix) (bool, error) {
	ipStr, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		ipStr = remoteAddr
	}

	ipAddr, err := netip.ParseAddr(ipStr)
	if err != nil {
		return false, err
	}

	for _, ipRange := range trustedIPs {
		if ipRange.Contains(ipAddr) {
			return true, nil
		}
	}

	return false, nil
}

// get the real IP from the proxy headers if present
func getRealIP(headers http.Header) string {
	var addr string

	for _, proxyHeader := range proxyIPHeaders {
		if value := headers.Get(proxyHeader); value != "" {
			addr = strings.SplitN(value, ",", 2)[0]
			break
		}
	}

	return addr
}

// get the scheme from the proxy headers if present
func getScheme(headers http.Header) string {
	var scheme string

	for _, schemaHeader := range schemeHeaders {
		if value := headers.Get(schemaHeader); value != "" {
			scheme = strings.ToLower(scheme)
			break
		}
	}

	return scheme
}
