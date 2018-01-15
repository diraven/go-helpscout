package helpscout

import "time"

type CustomerAddress struct {
	Id         int
	Lines      []string
	City       string
	State      string
	PostalCode string
	Country    string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
