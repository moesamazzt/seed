// Code generated by goctl. DO NOT EDIT.
package types

type ResultProductGetResp struct {
	Code    int32          `json:"code"`
	Message string         `json:"message"`
	Data    ProductGetResp `json:"data"`
}

type ResultProductListResp struct {
	Code    int32           `json:"code"`
	Message string          `json:"message"`
	Data    ProductListResp `json:"data"`
}

type ResultProductCreateResp struct {
	Code    int32             `json:"code"`
	Message string            `json:"message"`
	Data    ProductCreateResp `json:"data"`
}

type ProductGetReq struct {
	ProductId int64 `json:"productId"`
}

type ProductGetResp struct {
	ProductName     string `json:"productName"`
	ProductQuantity int64  `json:"productQuantity"`
	ProductEan      string `json:"productEan"`
}

type ProductListReq struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"pageSize"`
}

type ProductListResp struct {
	Products []ProductGetResp `json:"products"`
	Total    int32            `json:"total"`
}

type ProductCreateReq struct {
	ProductName     string `json:"productName"`
	ProductQuantity int64  `json:"productQuantity"`
	ProductEan      string `json:"productEan"`
}

type ProductCreateResp struct {
	ProductId int64 `json:"productId"`
}