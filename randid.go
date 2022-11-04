package randid

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
)

type ID [16]byte

func FromString(s string) (ID, error) {
	var id ID
	_, err := hex.Decode(id[:], []byte(s))
	return id, err
}

func (id ID) String() string {
	var s [32]byte
	hex.Encode(s[:], id[:])
	return string(s[:])
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*id, err = FromString(s)
	return err
}

func (id ID) Value() (driver.Value, error) {
	return id.String(), nil
}

func (id *ID) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
		*id = ID{}
	case []byte:
		if len(v) == 0 {
			*id = ID{}
			return nil
		}

		if len(v) != 16 {
			return id.Scan(string(v))
		}

		copy((*id)[:], v)
	case string:
		if len(v) == 0 {
			*id = ID{}
			return nil
		}

		i, err := FromString(v)
		if err != nil {
			return err
		}

		*id = i
		return nil
	}

	return nil
}
