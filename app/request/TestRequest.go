package request

import "time"

type TestAddRequest struct {
	Title     string     `json:"title" form:"title" validate:"required,max=255"`
	Content   string     `json:"content" form:"content" validate:"required"`
	Sorce     int        `json:"sorce" form:"sorce" validate:"required,number,min=0,max=100"`
	Category  string     `json:"category" form:"category" validate:"required,oneof=apple samsang oppo"`
	CreatedAt *time.Time `json:"created_at" form:"created_at"` // 非必传，不传默认为当前时间
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at"` // 非必传，不传默认为当前时间
}
type TestUpdateRequest struct {
	IdRequest
	Title    string `json:"title" form:"title" validate:"required,max=255"`
	Content  string `json:"content" form:"content" validate:"required"`
	Sorce    int    `json:"sorce" form:"sorce" validate:"required,number,min=0,max=100"`
	Category string `json:"category" form:"category" validate:"required,oneof=apple samsang oppo"`
}
