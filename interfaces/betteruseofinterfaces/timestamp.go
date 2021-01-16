package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time // Anonymous value
}

func (ts timestamp) String() string { // Notice that we are not using ts.Time.IsZero, this is because Time is an anonymous type
	// So therefore there is no need to get the ts.Time
	// If we add a name to time.Time then we would need to do something like ->>> ts.<name>.IsZero
	if ts.IsZero() {
		return "unknown"
	}
	const layout = "2006/01"
	return ts.Format(layout)
}

func toTimestamp(v interface{}) timestamp { // Any value either string or int returns a timestamp
	var ts timestamp
	if v == nil {
		return ts
	}
	var t int
	switch v := v.(type) {
	case int: // If the type is an int
		t = v
	case string: // If the type is an string
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}

func (ts timestamp) MarshalJSON() (data []byte, _ error) {
	// timestamp into integer
	return strconv.AppendInt(data, ts.Unix(), 10), nil
}

func (ts *timestamp) UnmarshalJSON(data []byte) error {
	*ts = toTimestamp(string(data))
	return nil
}
