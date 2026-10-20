package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

func (m MediaBrowseImage) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *MediaBrowseImage) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &m)
}

func (m MediaBrowseVideo) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *MediaBrowseVideo) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &m)
}

func (m MediaBrowseDocument) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *MediaBrowseDocument) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &m)
}
