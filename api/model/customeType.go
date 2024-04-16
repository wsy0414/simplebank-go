package model

import "time"

type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {

	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}
