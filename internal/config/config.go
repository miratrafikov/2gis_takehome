package config

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}
