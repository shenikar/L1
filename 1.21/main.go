package main

import (
	"fmt"
)

// Целевой интерфейс, который ожидает клиентский код
type Logger interface {
	Info(msg string)
	Error(msg string)
}

// Клиентский код, использующий наш интерфейс логгера
type App struct {
	logger Logger
}

func (a *App) Run() {
	a.logger.Info("Приложение запущено")

	a.logger.Error("Произошла ошибка")
}

// Сторонний логгер у которого несовместимый интерфейс
type ExternalLogger struct{}

func (l *ExternalLogger) LogInfo(msg string) {
	fmt.Println("[INFO]", msg)
}

func (l *ExternalLogger) LogError(msg string) {
	fmt.Println("[ERROR]", msg)
}

// Adapter для стороннего логгера
type ExternalLoggerAdapter struct {
	external *ExternalLogger
}

// Реализация нашего интерфейса через вызов методов стороннего логера
func (a *ExternalLoggerAdapter) Info(msg string) {
	a.external.LogInfo(msg)
}

func (a *ExternalLoggerAdapter) Error(msg string) {
	a.external.LogError(msg)
}

func main() {
	external := &ExternalLogger{}
	adapter := &ExternalLoggerAdapter{external: external}

	app := &App{logger: adapter}
	app.Run()
}
