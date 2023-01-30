package entity

type BasePage struct {
	PageSize int `form:"page_size"`
	PageNum  int `form:"page_num"`
}
