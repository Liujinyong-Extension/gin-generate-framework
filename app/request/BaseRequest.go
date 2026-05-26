package request

/*
*
分页请求参数
*/
type PageRequest struct {
	PageNum  int `form:"page_num" validate:"required,number,min=1"`
	PageSize int `form:"page_size" validate:"required,number,min=1"`
}

/*
*
id请求参数
*/
type IdRequest struct {
	Id int `form:"id" validate:"required,number,min=1"`
}
