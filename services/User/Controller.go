package user

import (
	"log"
	"main/interfaces"
	"main/routers"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	router         *routers.Router //Роутер, управляющий контроллером
	url_controller string          //часть полного url через который будет доступен контроллер
}

// Создание единственного экземпляра контроллера
var instanceController interfaces.Controller

// Метод, который запустит все обработчики контроллера с соответствующими параметрами
func (controller *UserController) RunHandler(url_pattern string) {
	//URL обработчиков(регистрация, вход, выход итд)
	// url_handlers = []url{"/reg", "/auth", "/exit", "/delete"}
	controller.reg(url_pattern + "/reg")
	controller.auth(url_pattern + "/auth")
	controller.exit(url_pattern + "/exit")
	controller.delete(url_pattern + "/delete")
}

func InitController(router *routers.Router, url_controller string) interfaces.Controller {
	if instanceController == nil {
		instanceController = &UserController{
			router:         router,
			url_controller: url_controller,
		}
		log.Println("Пользовательский контроллер создан!")

		//Запуск обработчиков
		instanceController.RunHandler(url_controller)

		return instanceController

	} else {
		log.Println("Пользовательский контроллер уже существует((")
	}
	return instanceController
}

func (controller *UserController) reg(full_url string) {
	//Приложение для которого инициализируется обработчик Reg контроллера
	app := controller.router.GetApp()

	//Регситрация пользователя
	app.Post(full_url, func(c *fiber.Ctx) error {
		return c.SendString("...")
	})
}

func (controller *UserController) auth(full_url string) {
	// app := controller.router.GetApp()
	// return
}

func (controller *UserController) exit(full_url string) {
	app := controller.router.GetApp()

	app.Get(full_url, func(request *fiber.Ctx) error { //Пока не ясно как будет работать(Заглушка)
		return request.SendString("")
	})
}

func (controller *UserController) delete(full_url string) {
	app := controller.router.GetApp()

	app.Delete(full_url, func(request *fiber.Ctx) error { //Пока не ясно как будет работать(Заглушка)
		return request.SendString("")
	})

}
