package model

type CarouselRequest struct {
}

type CarouselResponse struct {
	Items []CarouselItem `form:"items" json:"items"`
}

type CarouselItem struct {
	Id        uint   `form:"id" json:"id"`
	ImgPath   string `form:"img_path" json:"img_path"`
	ProductId uint   `form:"product_id" json:"product_id"`
	CreateAt  int64  `form:"create_at" json:"create_at"`
}
