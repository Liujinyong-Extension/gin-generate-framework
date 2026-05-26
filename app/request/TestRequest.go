package request

type AddRequest struct {
	Title    string `form:"title" validate:"required,max=255"`
	Content  string `form:"content" validate:"required"`
	Score    int    `form:"score" validate:"required,number,min=0,max=100"`
	Category string `form:"category" validate:"required,oneof=apple samsang oppo"`
}
