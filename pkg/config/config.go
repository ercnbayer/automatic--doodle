package config

import (
	"automatic-doodle/pkg/env"
	"automatic-doodle/types"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

var (
	configModule     ConfigModule
	configModuleOnce sync.Once
)

func New(
	log Logger,
) *ConfigModule {
	if log == nil {
		fmt.Println("Logger is not provided for ConfigModule!")
		os.Exit(1)
	}
	fmt.Println("config works")

	configModuleOnce.Do(func() {
		configModule = ConfigModule{
			log:   log,
			GoEnv: env.GO_ENV,
		}
	})

	return &configModule
}

type ConfigModule struct {
	log   Logger
	GoEnv types.GoEnv
}

func (mod *ConfigModule) GetConfig(key string) (string, error) {
	dest := os.Getenv(key)

	if len(dest) == 0 {
		mod.log.Warning("env key not found %s", key)
		return dest, errors.New("env key not found")
	}

	return dest, nil
}

func (mod *ConfigModule) GetConfigInt(key string) (int, error) {
	val := os.Getenv(key)

	if len(val) == 0 {
		mod.log.Warning("env key not found %s", key)
		return -1, errors.New("env key not found")
	}

	converted, err := strconv.Atoi(val)
	if err != nil {
		mod.log.Error("Number convertion failed for env key: %s! Error: %v", key, err)
		return -1, err
	}

	return converted, nil
}
