package routers

import (
	"main/interfaces"

	"github.com/gofiber/fiber/v2"
)

// Интерфейс контроллера
type Controller = func(*Router, string) interfaces.Controller

// Кастомный тип данных
type url = string

type Router struct {
	app      *fiber.App
	url_base url
}

func InitRouter(app *fiber.App, url_base string) *Router {
	return &Router{
		app:      app,
		url_base: url_base,
	}
}

// Настройка маршрутов и запуск обработчиков контроллеров
func (router *Router) Config(url_patterns map[string]func(*Router, string) interfaces.Controller) {
	for value_url, controller := range url_patterns {
		controller(router, value_url)
	}
}

// Получение доступа к app в Router из вне
func (router *Router) GetApp() *fiber.App {
	return router.app
}
