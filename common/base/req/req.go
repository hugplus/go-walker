package req

type BaseReq struct {
	ReqId string `json:"reqId" form:"reqId"` // 请求id 链路跟踪
}

type IdReq struct {
	Id int `json:"id" form:"id"` // 主键ID
}

type StrIdReq struct {
	Id string `json:"id" form:"id"` // 主键ID
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"` //多id
}

type PageInfo struct {
	Page    int    `json:"page" form:"page"`       // 页码
	Size    int    `json:"size" form:"size"`       // 每页大小
	Keyword string `json:"keyword" form:"keyword"` //关键字
}