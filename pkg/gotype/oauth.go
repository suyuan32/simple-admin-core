package gotype

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

type OAuthStyle uint8

func (s *OAuthStyle) Value() (driver.Value, error) {
	return int64(*s), nil
}

func (s *OAuthStyle) Scan(v any) error {
	s2 := asString(v)
	parseUint, err := strconv.ParseUint(s2, 10, 8)
	if err != nil {
		return fmt.Errorf("invalid database type: %T %v", v, v)
	}
	*s = OAuthStyle(parseUint)
	return nil
}

const (
	OAuthStyleAuto OAuthStyle = iota
	OAuthStyleThirdParty
	OAuthStyleUsernamePassword
)
