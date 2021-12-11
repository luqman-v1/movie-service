package util

import "google.golang.org/grpc/status"

func NotFoundError(model string) error {
	return status.Errorf(NotFound, "The "+model+" resounce not found")
}
