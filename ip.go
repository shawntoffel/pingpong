package pingpong

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func findOriginatingIP(r *http.Request) (string, error) {
	xForwardedFor := r.Header.Get("x-forwarded-for")
	if xForwardedFor == "" {
		remoteAddr, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			return "", fmt.Errorf("failed to split host and port from remote address")
		}

		return remoteAddr, nil
	}

	xForwardedFor = strings.Split(xForwardedFor, ",")[0]
	xForwardedFor = strings.TrimSpace(xForwardedFor)

	ip := net.ParseIP(xForwardedFor)
	if ip == nil {
		return "", fmt.Errorf("X-Forwarded-For header did not contain a valid IP")
	}

	return ip.String(), nil
}

// OriginatingIP responds with the originating IP address of a client
func OriginatingIP(w http.ResponseWriter, r *http.Request) {
	originatingIP, err := findOriginatingIP(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, originatingIP)
}
