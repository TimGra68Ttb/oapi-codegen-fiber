package handlers

import "github.com/timgra68ttb/oapi-codegen-fiber/controller"

func (h *MyHandler) CheckHandler() controller.ICheckController {
	checkController := controller.NewCheckController()
	return checkController
}
