package utils

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
