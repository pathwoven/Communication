package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	DBUser     string `yaml:"user"`
	DBPassword string `yaml:"password"`
	DBHost     string `yaml:"host"`
	DBPort     string `yaml:"port"`
	DBName     string `yaml:"database"`
	DBCharset  string `yaml:"charset"`
}
type ServerConfig struct {
	ServerPort string `yaml:"port"`
}
type SecurityConfig struct {
	SessionConfig `yaml:"session"`
}
type SessionConfig struct {
	HashKey  string `yaml:"hash-key"`
	BlockKey string `yaml:"block-key"`
}
type Config struct {
	DatabaseConfig `yaml:"mysql"`
	ServerConfig   `yaml:"server"`
	SecurityConfig `yaml:"security"`
}

var AppConfig Config

func InitConfig() {
	// 加载config.yml文件
	data, err := os.ReadFile(filepath.Join("config", "config.yml"))
	if err != nil {
		log.Fatal("无法读取配置文件: ", err)
	}
	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatal("配置文件解析失败: ", err)
	}
}
