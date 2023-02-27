package uuidx

import (
	"github.com/gofrs/uuid/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

// NewUUID returns a new UUID.
func NewUUID() uuid.UUID {
	id, err := uuid.NewV7()
	if err != nil {
		logx.Errorw("fail to generate UUID", logx.Field("detail", err))
		return uuid.UUID{}
	}
	return id
}

// ParseUUIDSlice parses the UUID string slice to UUID slice.
func ParseUUIDSlice(ids []string) []uuid.UUID {
	var result []uuid.UUID
	for _, v := range ids {
		p, err := uuid.FromString(v)
		if err != nil {
			logx.Errorw("fail to parse string to UUID", logx.Field("detail", err))
			return nil
		}
		result = append(result, p)
	}
	return result
}

// ParseUUIDString parses UUID string to UUID type.
func ParseUUIDString(id string) uuid.UUID {
	result, err := uuid.FromString(id)
	if err != nil {
		logx.Errorw("fail to parse string to UUID", logx.Field("detail", err))
		return uuid.UUID{}
	}
	return result
}
