package models

import (
	"fmt"
	"strconv"
	"time"
)

var (
	NilUnixTime = UnixTime(time.Time{})
)

type UnixTime time.Time

func UnixTimeNow() UnixTime {
	now := time.Now().Unix()
	nowUnix := UnixTime(time.Unix(now, 0))
	return nowUnix
}

func (t *UnixTime) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

func (t *UnixTime) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = UnixTime(time.Unix(int64(ts), 0))

	return nil
}

func (t *UnixTime) String() string {
	tm := time.Time(*t).UTC().In(time.Local)
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
}

func (t *UnixTime) StringRaw() string {
	return fmt.Sprintf("%d", time.Time(*t).UnixNano())
}

func (t *UnixTime) Equal(t2 UnixTime) bool {
	return time.Time(*t).Equal(time.Time(t2))
}

func (t *UnixTime) After(t2 UnixTime) bool {
	return time.Time(*t).After(time.Time(t2))
}
