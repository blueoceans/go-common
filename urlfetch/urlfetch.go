// +build appengine

package urlfetch

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"appengine"
	"appengine/urlfetch"
	"github.com/blueoceans/go-common/errors"
	u "github.com/blueoceans/go-common/logutil"
)

const (
	reasonDeadlineExceededTimeout = "Deadline exceeded (timeout)"

	reasonUrlfetchClosed               = "urlfetch: CLOSED"
	reasonUrlfetchConnectionErr        = "urlfetch: CONNECTION_ERROR"
	reasonUrlfetchDeadlineExceeded     = "urlfetch: DEADLINE_EXCEEDED"
	reasonUrlfetchFetchErr             = "urlfetch: FETCH_ERROR"
	reasonUrlfetchInternalTransientErr = "urlfetch: INTERNAL_TRANSIENT_ERROR"

	sec1 = 1 * time.Second
)

var (
	errNoWait = []string{
		reasonDeadlineExceededTimeout,
		reasonUrlfetchClosed,
		reasonUrlfetchConnectionErr,
		reasonUrlfetchDeadlineExceeded,
		reasonUrlfetchFetchErr,
	}
	errWait = []string{
		reasonUrlfetchInternalTransientErr,
	}
)

func ClientGet(
	c appengine.Context,
	url string,
) (
	*http.Response,
	error,
) {
	return clientDo(
		c,
		func() (
			*http.Response,
			error,
		) {
			return urlfetch.Client(c).Get(url)
		},
	)
}

func ClientPost(
	c appengine.Context,
	url string,
	mimeType string,
	body io.Reader,
) (
	*http.Response,
	error,
) {
	return clientDo(
		c,
		func() (
			*http.Response,
			error,
		) {
			return urlfetch.Client(c).Post(url, mimeType, body)
		},
	)
}

func ClientPostForm(
	c appengine.Context,
	url string,
	data url.Values,
) (
	*http.Response,
	error,
) {
	return clientDo(
		c,
		func() (
			*http.Response,
			error,
		) {
			return urlfetch.Client(c).PostForm(url, data)
		},
	)
}

func clientDo(
	c appengine.Context,
	do func() (*http.Response, error),
) (
	*http.Response,
	error,
) {
retry:
	response, err := do()
	switch {
	case err == nil:
		return response, nil
	case IsErrNoWait(err):
		u.Infof(c, "%q", err)
		goto retry
	case IsErrWait(err):
		u.Infof(c, "%q", err)
		time.Sleep(sec1)
		goto retry
	default:
		return nil, err
	}
}

// IsErrNoWait returns is whether it can re-try soon or not.
func IsErrNoWait(
	err error,
) bool {
	return errors.ContainsErrorMessage(err, errNoWait)
}

// IsErrWait returns is whether it must wait for next re-try or not.
func IsErrWait(
	err error,
) bool {
	return errors.ContainsErrorMessage(err, errWait)
}
