package config

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	Showline     bool   `yaml:"showline"`
	LogInConsole bool   `yaml:"log_in_console"`
}
