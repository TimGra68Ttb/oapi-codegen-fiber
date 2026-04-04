package handlers

import "github.com/timgra68ttb/oapi-codegen-fiber/controller"

func (f *Factory) CheckController() controller.ICheckController {
	checkController := controller.NewCheckController()
	return checkController
}
