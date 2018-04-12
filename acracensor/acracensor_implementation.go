package acracensor

type AcraCensor struct {
	handlers []QueryHandlerInterface

}

func (acraCensor *AcraCensor) AddHandler(handler QueryHandlerInterface) {
	acraCensor.handlers = append(acraCensor.handlers, handler)
}

func (acraCensor *AcraCensor) RemoveHandler(handler QueryHandlerInterface) {
	for index, handlerFromRange := range acraCensor.handlers {
		if handlerFromRange == handler {
			acraCensor.handlers = append(acraCensor.handlers[:index], acraCensor.handlers[index+1:]...)
		}
	}
}

func (acraCensor *AcraCensor) HandleQuery(query string) error {
	for _, handler := range acraCensor.handlers {
		if err := handler.CheckQuery(query); err != nil {
			return err
		}
	}
	return nil
}
