package cmd

import (
	"os"

	"github.com/LightAir/cabty/internal/app"
	"github.com/LightAir/cabty/internal/config"
	"github.com/LightAir/cabty/internal/packer"
	"github.com/LightAir/cabty/internal/uploader"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const ErrorExitCode = 1

var rootCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "cabty",
	Short: "Tools for crypt and backup files to you Yandex.Disk",
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out: os.Stderr,
			},
		)

		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

		return nil
	},
	Run: func(_ *cobra.Command, _ []string) {
		err := godotenv.Load()
		if err != nil {
			log.Err(err).Msg("Error loading .env file")

			os.Exit(ErrorExitCode)
		}

		var cfg config.AppConfig

		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Err(err).Msg("Error read env")

			os.Exit(ErrorExitCode)
		}

		log.Info().Msg("programm started")

		packerService := packer.NewService()

		uploaderService := uploader.NewService()

		service := app.NewService(packerService, uploaderService)

		if err = service.Run(&cfg); err != nil {
			log.Err(err).Msg("Error read env")

			return
		}

		log.Info().Msg("programm completed")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(ErrorExitCode)
	}
}
