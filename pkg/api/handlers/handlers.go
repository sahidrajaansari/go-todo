package handlers

type Handlers struct {
	TodoHandler *TodoHandler
}

func NewHandler(th *TodoHandler) *Handlers {
	return &Handlers{
		TodoHandler: th,
	}
}
