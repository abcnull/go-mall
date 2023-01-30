package model

type AddressListRequest struct {
	Keyword  string `json:"keyword"`
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
}

type AddressListResponse struct {
	Addresses []Address `json:"addresses"`
	Total     int64     `json:"total"`
}

type Address struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type AddressDetailRequest struct {
}

type AddressDetailResponse struct {
	Address Address `json:"address"`
}

type CreateAddressRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type CreateAddressResponse struct {
	Address Address `json:"address"`
}

type UpdateAddressRequest struct {
	Address Address `json:"address"`
}

type UpdateAddressResponse struct {
	Address Address `json:"address"`
}

type DeleteAddressRequest struct {
	Address []Address `json:"address"`
}

type DeleteAddressResponse struct {
}
