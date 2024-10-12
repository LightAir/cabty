package config

type AppConfig struct {
	BaseDir      string `env:"BASE_DIR"`
	YandexConfig YandexConfig
}

type YandexConfig struct {
	OAuthToken string `env:"YANDEX_OAUTH_TOKEN"`
	SavePath   string `env:"YANDEX_SAVE_PATH"`
}
