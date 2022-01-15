package server

import (
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

type Config struct{
	Port         string `yaml:"port"`
	PostgresLink string `yaml:"psql"`
}
func NewConfig(path string) (*Config, error){
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(f)
	if err != nil{
		return nil, err
	}
	var res Config
	if err := yaml.Unmarshal(data, &res); err != nil{
		return nil, err
	}
	return &res, nil
}