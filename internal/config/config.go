package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string              `yaml:"env" env-default:"local"`
	StoragePath string              `yaml:"storage_path" env-required:"true"`
	TokenTtl    time.Duration       `yaml:"token_ttl" env-default:"3h"`
	GRPC        GRPCConfig          `yaml:"grpc"`
	PSQL        PSQL                `yaml:"psql"`
	Redis       Redis               `yaml:"redis"`
	Modules     ModulesEnableConfig `yaml:"modules"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-default:"44044"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

type ModulesEnableConfig struct {
	DatabaseEnabled int `yaml:"database_enabled" env-default:"1"`
	RedisEnabled    int `yaml:"redis_enabled" env-default:"0"`
}

type PSQL struct {
	Username string `yaml:"user" env-required:"true"`
	Dbname   string `yaml:"dbname" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port"  env-default:"5432"`
}

type Redis struct {
	Addr     string `yaml:"addr" env-required:"true"`
	Password string `yaml:"password" env-default:""`
	Db       int    `yaml:"db" env-default:"0"`
}

// Эта функция загружает конфигурацию и завершает программу при ошибке
func MustLoadConfig() *Config {

	//Проверяется аргумент командной строки -config.
	//Если не указан, берётся значение переменной окружения CONFIG_PATH.
	//Если путь так и не найден, в path остаётся пустая строка.
	path := func() string {
		var res string

		flag.StringVar(&res, "config", "", "path to config .yaml file")
		flag.Parse()

		if res == "" {
			res = os.Getenv("CONFIG_PATH")
		}

		return res
	}()

	//Если путь пустой → panic("config path is empty").
	//Если файл не найден → panic("config path does not exist ...").
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic(fmt.Sprintf("config path does not exists %s", path))
	}

	//cleanenv.ReadConfig(path, &cfg) читает YAML-файл и записывает данные в cfg.
	//Если чтение не удалось, программа падает с panic.
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(fmt.Sprintf("failed while reading config %s", err.Error()))
	}

	return &cfg
}
