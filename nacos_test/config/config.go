package config

type UserSrvConfig struct {
	Host string `mapstucture:"host" json:"host"`
	Port int    `mapstucture:"port" json:"port"`
	Name string `mapstucture:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type AliSMSConfig struct {
	ApiKey    string `mapstructure:"key" json:"key"`
	ApiSecret string `mapstructure:"secret" json:"secret"`
}

type ConsulConfig struct {
	Host string `mapstucture:"host" json:"host"`
	Port int    `mapstucture:"port" json:"port"`
}

type RedisConfig struct {
	Host   string `mapstucture:"host" json:"host"`
	Port   int    `mapstucture:"port" json:"port"`
	Expire int    `mapstucture:"expire" json:"expire"`
}

type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	Port        int           `mapstructure:"port" json:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	AliSMSInfo  AliSMSConfig  `mapstructure:"sms" json:"sms"`
	RedisInfo   RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
}
