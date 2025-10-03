package data

import "time"

type Operation struct {
	Op  string
	Val []byte
	Ttl time.Time
}
