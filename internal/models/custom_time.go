package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b[1 : len(b)-1]) // remove quotes
	t, err := time.Parse(ctLayout, str)
	if err != nil {
		return fmt.Errorf("error parsing time: %w", err)
	}
	ct.Time = t
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", ct.Format(ctLayout))), nil
}

func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Format(ctLayout), nil
}

func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*ct = CustomTime{Time: time.Time{}}
		return nil
	}
	t, err := time.Parse(ctLayout, value.(string))
	if err != nil {
		return fmt.Errorf("error scanning time: %w", err)
	}
	*ct = CustomTime{Time: t}
	return nil
}

func (ct CustomTime) String() string {
	return ct.Format(ctLayout)
}