package packer

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Packer struct{}

func NewService() *Packer {
	return &Packer{}
}

func (ps *Packer) Pack(_ context.Context, _ string) (string, error) {
	log.Info().Msg("pack in progress...")

	return "", nil
}
