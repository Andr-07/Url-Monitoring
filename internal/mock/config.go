package mock

import configs "go-monitoring/config"

type MockConfig struct {
	AuthSecret string
}

func NewMockConfig() *configs.Config {
	return &configs.Config{
		Auth: configs.AuthConfig{
			Secret: "secret",
		},
	}
}
