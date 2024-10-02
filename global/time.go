package global

import (
	"database/sql/driver"
	"fmt"
	"go-gin/cons"
	"time"
)

// 全局定义
type Time time.Time

// MarshalJSON 实现 json.Marshaler 接口
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(cons.DateTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, cons.DateTimeFormat)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+cons.DateTimeFormat+`"`, string(data), time.Local)
	if err != nil {
		return err
	}
	*t = Time(now)
	return nil
}

// String 实现 Stringer 接口
func (t Time) String() string {
	return time.Time(t).Format(cons.DateTimeFormat)
}

// local 返回本地时间
func (t Time) local() time.Time {
	loc, _ := time.LoadLocation(cons.TIMEZONE)
	return time.Time(t).In(loc)
}

// Value 实现 driver.Valuer 接口
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan 实现 sql.Scanner 接口
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("不能将 %v 转换成时间戳", v)
}
