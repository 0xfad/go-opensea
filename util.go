package goopensea

import (
	"encoding/json"
	"strings"
	"time"
)

type CreationDate time.Time

func (cd *CreationDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}

	*cd = CreationDate(t)

	return nil
}

func (cd *CreationDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*cd))
}
