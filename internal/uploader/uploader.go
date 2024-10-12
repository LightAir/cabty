package uploader

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Uploader struct{}

func NewService() *Uploader {
	return &Uploader{}
}

func (us *Uploader) Upload(_ context.Context, _ string) error {
	log.Info().Msg("upload in progress...")

	return nil
}
