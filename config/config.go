package config

type ServerConfig struct {
	Name               string             `mapstructure:"name" json:"name"`
	Host               string             `mapstructure:"host" json:"host"`
	Port               int64              `mapstructure:"port" json:"port"`
	Tags               []string           `mapstructure:"tags" json:"tags"`
	JWTConfig          JWTConfig          `mapstructure:"jwt" json:"jwt"`
	UserServerConfig   UserServerConfig   `mapstructure:"user-server" json:"user-server"`
	PublicServerConfig PublicServerConfig `mapstructure:"public-server" json:"public-server"`
}

type JWTConfig struct {
	Issuer     string `mapstructure:"issuer" json:"issuer"`
	Expire     int64  `mapstructure:"expire" json:"expire"`
	SigningKey string `mapstructure:"key" json:"signingKey"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type PublicServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
