package developer

import (
    "time"
)

type Developer struct {
    Email		string
    Pass		[]byte
    Auth 		time.Time
    Address		string
    Balance		float64
}
