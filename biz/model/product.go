package model

type CreateProductRequest struct {
	Name       string `form:"name" json:"name"`               // 商品名称
	CategoryId uint   `form:"category_id" json:"category_id"` // 品类 id

	Title   string `form:"title" json:"title"`       // 商品标题
	Info    string `form:"info" json:"info"`         // 商品信息
	ImgPath string `form:"img_path" json:"img_path"` // 商品图片 todo: 尝试一下注释掉也可以？
	Locate  string `form:"locate" json:"locate"`     // 地址 // todo: 这里商品表中待添加

	//ProductDetail ProductDetail `form:"product_detail" json:"product_detail"` // 商品细节信息 todo: 这个以后补上

	OriginPrice   string      `form:"origin_price" json:"origin_price"`     // 原价
	DiscountPrice string      `form:"discount_price" json:"discount_price"` // 售卖价格
	Freight       FreightType `form:"freight" json:"freight"`               // 运输方式 // todo: 待添加
}

// ProductDetail 商品详细信息
type ProductDetail struct {
	/* 通用 */
	BrandName       string `form:"brand_name" json:"brand_name"`             // 品牌名称
	Condition       string `form:"condition" json:"condition"`               // 新旧程度
	PurchaseChannel string `form:"purchase_channel" json:"purchase_channel"` // 购买渠道

	/* 较通用 */
	FunctionStatus string `form:"function_status" json:"function_status"` // 功能状态

	/* 保温杯 */
	Capability string `form:"capability" json:"capability"` // 容量
	KeepWarm   string `form:"keep_warm" json:"keep_warm"`   // 是否保温

	/* 衣物 */
	ClotheSize string `form:"clothe_size" json:"clothe_size"` // 尺码
	FitSeason  string `form:"fit_season" json:"fit_season"`   // 适用季节
	Sleeves    string `form:"sleeves" json:"sleeves"`         // 袖长短袖 or 长袖

	/* 包 */
	Style string `form:"style" json:"style"` // 款式
}

type CreateProductResponse struct {
	Id         uint   `json:"id"`          // 商品 id
	Name       string `json:"name"`        // 商品名称
	CategoryId uint   `json:"category_id"` // 商品品类 id

	Tile    string `json:"title"`                    // 商品标题
	Info    string `json:"info"`                     // 商品描述
	ImgPath string `form:"img_path" json:"img_path"` // 商品图片
	Locate  string `form:"locate" json:"locate"`     // 地址 // todo: 这里商品表中待添加

	OriginPrice   string      `form:"origin_price" json:"origin_price"`     // 原价
	DiscountPrice string      `form:"discount_price" json:"discount_price"` // 售卖价格
	Freight       FreightType `form:"freight" json:"freight"`               // 运输方式 // todo: 待添加

	OnSale bool `json:"on_sale"` // 是否上架

	BossId     uint   `json:"boss_id"`
	BossName   string `json:"boss_name"`
	BossAvatar string `json:"boss_avatar"`

	ExposureTimes      int64 `json:"exposure_times"`      // 曝光量
	ClickTimes         int64 `json:"click_times"`         // 点击量
	CommunicationTimes int64 `json:"communication_times"` // 沟通数

	CreateAt int64 `json:"create_at"`
	UpdateAt int64 `json:"update_at"`
}

// ProductListRequest 商品列表
type ProductListRequest struct {
	Keyword  string `form:"keyword" json:"keyword"`
	PageSize int    `form:"page_size" json:"page_size,required"`
	PageNum  int    `form:"page_num" json:"page_num,required"`
}

type ProductListResponse struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
}

type Product struct {
	Id         uint `json:"id"`
	CategoryId uint `json:"category_id"`

	Title         string      `json:"title"`
	Info          string      `json:"info"`
	CoverImgPath  string      `json:"cover_img_path"`
	AllImgPath    []string    `json:"all_img_path"`
	OriginPrice   float64     `json:"origin_price"`
	DiscountPrice float64     `json:"discount_price"`
	Locate        string      `json:"locate"`
	Freight       FreightType `json:"freight"`

	BossId     uint   `json:"boss_id"`
	BossName   string `json:"boss_name"`
	BossAvatar string `json:"boss_avatar"`

	ExposureTimes      int64 `json:"exposure_times"`
	ClickTimes         int64 `json:"click_times"`
	CommunicationTimes int64 `json:"communication_times"`
}

// ProductDetailRequest 查询详细商品信息
type ProductDetailRequest struct {
}

type ProductDetailResponse struct {
	Product Product `json:"product"`
}

type AutoGainCategoryRequest struct {
	Title   string `form:"title" json:"title"`
	Info    string `form:"info" json:"info"`
	ImgPath string `form:"img_path" json:"img_path"` // todo: 注释掉也可以？
}

type AutoGainCategoryResponse struct {
	Id   uint   `form:"id" json:"id"`     // 种类 id
	Name string `form:"name" json:"name"` // 种类名称
}
