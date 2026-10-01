package domain

import "time"

type Page struct {
	URL  string
	Time time.Time
	Body []byte
}
