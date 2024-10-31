package config

import "fmt"

type Config struct {
	Port           int
	WebFilesFolder string
	TemplateFolder string
}

func NewConfig() *Config {
	webFilesFolder := "../../web"
	return &Config{
		Port:           8080,
		WebFilesFolder: webFilesFolder,
		TemplateFolder: fmt.Sprintf("%s/templates", webFilesFolder),
	}
}
