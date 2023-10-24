package common

import "time"

type SQLmodel struct {
	Id         int        `json:"id"`
	Created_at *time.Time `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
}
