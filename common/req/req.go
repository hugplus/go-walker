package req

type IdReq struct {
	Id int `json:"id" form:"id"` // 主键ID
}

type StrIdReq struct {
	Id string `json:"id" form:"id"` // 主键ID
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}
