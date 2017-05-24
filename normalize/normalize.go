package normalize

import (
	"net/url"
	"strings"
)

// URI removes the standard port number.
func URI(
	originURI string,
) (
	string,
	error,
) {
	uri, err := url.ParseRequestURI(originURI)
	if err != nil {
		return "", err
	}

	splits := strings.SplitN(uri.Host, ":", 2) // hostname:port

	// remove the standard port number
	switch len(splits) {
	case 2:
		switch {
		case
			uri.Scheme == "http" && splits[1] == "80",
			uri.Scheme == "https" && splits[1] == "443":
			uri.Host = splits[0]
		}
	}

	uri.User = nil
	return uri.String(), nil
}
