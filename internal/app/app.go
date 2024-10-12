package app

import (
	"context"

	"github.com/LightAir/cabty/internal/config"
	"github.com/cockroachdb/errors"
)

type Packer interface {
	Pack(ctx context.Context, packDir string) (string, error)
}

type Uploader interface {
	Upload(ctx context.Context, path string) error
}

type App struct {
	Packer   Packer
	Uploader Uploader
}

func NewService(packer Packer, uploader Uploader) *App {
	return &App{
		Packer:   packer,
		Uploader: uploader,
	}
}

func (a App) Run(cfg *config.AppConfig) error {
	ctx := context.Background()

	pipe, err := a.Packer.Pack(ctx, cfg.BaseDir)
	if err != nil {
		return errors.Wrap(err, "Run")
	}

	err = a.Uploader.Upload(ctx, pipe)
	if err != nil {
		return errors.Wrap(err, "Run")
	}

	return nil
}
