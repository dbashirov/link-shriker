package config

type Config struct {
	BaseURL                 string `env:"base_url,default=http://localhost:8080"`
	Host                    string `env:"host,default=0.0.0.0"`
	Port                    int    `env:"port,default=8080"`
	TelegramContactUsername string `env:"telegram_contact_username,default=tomakado"`
	DB                      DBConfig
	GitHub                  GitHubConfig
	Auth                    AuthConfig
}

type DBConfig struct {
	DSN      string `env:"mongodb_dsn"`
	Database string `env:"mongodb_database"`
}

type GitHubConfig struct {
	ClientID     string `env:"github_client_id"`
	ClientSecret string `env:"github_client_secret"`
}

type AuthConfig struct {
	JWTSecretKey     string `env:"jwt_secret_key"`
	AllowedGitHubOrg string `env:"allowed_github_org,default=defer-panic"`
}
