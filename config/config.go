package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type SourceRegistry struct {
	Name   string   `yaml:"name"`
	Images []string `yaml:"images"`
}

type Source struct {
	Registries []SourceRegistry `yaml:"registries"`
}

type TargetRepository struct {
	Name string `yaml:"name"`
}

type Target struct {
	Registry     string             `yaml:"registry"`
	Repositories []TargetRepository `yaml:"repositories"`
}

type Config struct {
	Source Source `yaml:"source"`
	Target Target `yaml:"target"`
}

func LoadConfig(filePath string) (*Config, error) {
	config := &Config{}
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
