package main

import (
	"log"
	"main/interfaces"
	"main/routers"
	"main/scripts"
	user "main/services/User"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Запуск фронтенда
	go scripts.RunClient()

	app := fiber.New()
	//Список контроллеров-сервисов
	url_patterns := map[string]func(*routers.Router, string) interfaces.Controller{
		"/user": user.InitController,
	}

	router := routers.InitRouter(app, "/")
	router.Config(url_patterns)

	log.Fatal(app.Listen(":3001"))

}
