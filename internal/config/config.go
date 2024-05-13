package config

type Config struct {
	Server  Server  `yaml:"server"`
	Logging Logging `yaml:"logging"`
}

type Server struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}

type Logging struct {
	Level string `yaml:"level"`
}
