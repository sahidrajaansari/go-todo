package todo

type UpdateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateTodoResponse struct {
	ID          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	UpdatedAt   string `json:"updatedat"`
}

func (req *UpdateTodoRequest) SetDefaultValues() {
	if req.Title == "" {
		req.Title = ""
	}
	if req.Description == "" {
		req.Description = ""
	}
	if req.Status == "" {
		req.Status = ""
	}
}
