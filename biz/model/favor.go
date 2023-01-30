package model

type FavorListRequest struct {
	Keyword  string `json:"keyword"`
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
}

type FavorListResponse struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
}

type AddFavorRequest struct {
	ProductId uint `json:"product_id"`
}

type AddFavorResponse struct {
	ProductId  uint `json:"product_id"`
	CategoryId uint `json:"category_id"`
	BossId     uint `json:"boss_id"`
}

type DeleteInFavorRequest struct {
	ProductId uint `json:"product_id"`
}

type DeleteInFavorResponse struct {
}
