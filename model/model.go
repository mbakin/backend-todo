package model

import "time"

type Todo struct {
	ID   int       `json:"id"`
	Todo string    `json:"todo"`
	Time time.Time `json:"time"`
}
