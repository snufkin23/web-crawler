package domain

import "errors"

var (
	ErrUnsupportedScheme = errors.New("unsupported URL scheme")
	ErrNonHTML           = errors.New("response is not HTML")
	ErrRobotsDisallowed  = errors.New("robots.txt disallows URL")
	ErrMaxDepth          = errors.New("max depth reached")
)
