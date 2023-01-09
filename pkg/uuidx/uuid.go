package uuidx

import (
	"github.com/gofrs/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

func NewUUID() uuid.UUID {
	id, err := uuid.NewV7()
	if err != nil {
		logx.Errorw("fail to generate UUID", logx.Field("detail", err))
		return uuid.UUID{}
	}
	return id
}

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

func ParseUUIDString(id string) uuid.UUID {
	result, err := uuid.FromString(id)
	if err != nil {
		logx.Errorw("fail to parse string to UUID", logx.Field("detail", err))
		return uuid.UUID{}
	}
	return result
}
