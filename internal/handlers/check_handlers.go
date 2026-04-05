package handlers

func (f *Factory) CheckController() ICheckController {
	checkController := NewCheckController()
	return checkController
}
