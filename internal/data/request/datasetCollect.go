package request

type DatasetsCollectReq struct {
	Token    string `form:"token" json:"token" binding:"required"`
	Sequence int    `form:"sequence" json:"sequence" binding:"required"`
}
