package echoext

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type Timestamp time.Time

// UnmarshalParam echo api @see https://echo.labstack.com/guide/request
func (t *Timestamp) UnmarshalParam(src string) error {
	if src != "" {
		m, err := strconv.ParseInt(src, 10, 64)
		if err != nil {
			return err
		}

		ts := time.Unix(0, m*int64(time.Millisecond))
		*t = Timestamp(ts)
	}
	return nil
}

// MarshalJSON echo api json response
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	if t != nil {
		ts := time.Time(*t)
		return []byte(fmt.Sprintf(`"%d"`, ts.UnixNano()/int64(time.Millisecond))), nil
	}
	return nil, nil
}

// for sql log, print readable format
func (t Timestamp) String() string {
	ts := time.Time(t)
	return ts.Format(time.RFC3339)
}

// insert into database conversion
func (t Timestamp) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// read from database conversion
func (t *Timestamp) Scan(src interface{}) error {
	if val, ok := src.(time.Time); ok {
		*t = Timestamp(val)
	}
	return nil
}
