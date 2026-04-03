package config

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	}
	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Db       int    `mapstructure:"db"`
	}
}
