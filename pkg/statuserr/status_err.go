package statuserr

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewInternalError(msg string) error {
	return status.Error(codes.Internal, msg)
}

func NewInvalidArgumentError(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}

func NewNotFoundError(msg string) error {
	return status.Error(codes.NotFound, msg)
}

func NewAlreadyExistsError(msg string) error {
	return status.Error(codes.AlreadyExists, msg)
}

func NewUnauthenticatedError(msg string) error {
	return status.Error(codes.Unauthenticated, msg)
}
