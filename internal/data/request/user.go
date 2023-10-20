package request

type UserReq struct {
	OpenId      string `form:"openId" json:"openId"`
	FullName    string `form:"fullName" json:"fullName" binding:"required"`
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" binding:"required"`
}

type UserListReq struct {
	//页面大小
	StartPage int `form:"start_page" json:"start_page"`
	//起始页
	PageSize int `form:"page_size" json:"page_size"`
}

type PermissionReq struct {
	Id   int `json:"id"`
	Role int `json:"role"`
}
