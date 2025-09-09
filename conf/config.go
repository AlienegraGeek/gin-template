package conf

var Config = struct {
	DB      DB     `yaml:"db"`
	Zap     Zap    `yaml:"zap"`
	GinPort string `yaml:"gin-port"`
}{}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
