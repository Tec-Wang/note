package diary

type idParam struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
