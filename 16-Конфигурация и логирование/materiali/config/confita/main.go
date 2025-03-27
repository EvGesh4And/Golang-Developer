package main

import (
	"context"
	"log/slog"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
)

// Config - структура для хранения конфигурации
// Поля могут быть загружены из разных источников (файл, переменные окружения и т.д.)
type Config struct {
	Port  int    `config:"port"`
	Db    string `config:"db" json:"db"`
	Debug bool   `config:"debug"`
}

func main() {
	// Создаем загрузчик конфигурации с несколькими backend'ами
	loader := confita.NewLoader(
		file.NewBackend("base.json"),   // Загрузка базовой конфигурации
		file.NewBackend("custom.json"), // Загрузка пользовательской конфигурации (переопределяет base.json)
		env.NewBackend(),               // Загрузка переменных окружения
		flags.NewBackend(),             // Загрузка аргументов командной строки
	)

	var cfg Config

	// Загружаем конфигурацию
	if err := loader.Load(context.Background(), &cfg); err != nil {
		slog.Error("Ошибка загрузки конфигурации", "error", err)
		panic(err)
	}

	// Вывод загруженной конфигурации
	slog.Info("Конфигурация загружена",
		"port", cfg.Port,
		"db", cfg.Db,
		"debug", cfg.Debug,
	)
}

// PORT=12 go run main.go -port=123
