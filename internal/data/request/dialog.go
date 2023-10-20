package request

type DialogReq struct {
	// 发音人
	Speaker string `form:"speaker" json:"speaker"  binding:"required"`
	// 速度
	Speed int `form:"speed" json:"speed"`
	// 音量
	Volume int `form:"volume" json:"volume"`
	// 1:low   2:medium  3:high  默认2:medium
	Pitch int `form:"pitch" json:"pitch"`
	// 0:达闼  1：Google  2：科大讯飞  3:CPAll
	ASRServiceProvider int `form:"asrType" json:"asrType"`
	//语言类型：    0:CH  1:EN 2:TCH 3:JA 4:ES 5:TH
	LangType int `form:"langType" json:"langType"`
	//服务类型： 0:通用  1:数码通讯  2：医疗康养
	ServiceType int `form:"serviceType" json:"serviceType"`
	// NLP agent
	NlpAgent int `form:"nlpAgent" json:"nlpAgent" binding:"required"`
	// 环境类型：0：default 生成环境 1:研发环境
	EnvType int `form:"envType" json:"envType" binding:"required"`
	// 机器人所在位置。格式：经度;纬度
	Position string `form:"position" json:"position" binding:"required"`
	// 交互方式  0:文本    1：语音    2：监听   3：保持心跳
	InteractionType int `form:"interactionType" json:"interactionType" binding:"required"`
	// 交互为文本时 的文本
	Text string `form:"text" json:"text"`
	// sessionID
	SessionID string `form:"sessionID" json:"sessionID"`
	// VAD超时时间 ms  0--极速模式（默认300ms）1--交互模式（600ms）2--听写模式（1200ms）
	VADPauseMode int `json:"vadPauseMode"`
}

var VendorMap = map[int]string{
	0: "CloudMinds",
	1: "Google",
	2: "IFlyTek",
	3: "CPAll",
}

var ServiceTypeMap = map[int]string{
	0: "Common",
	1: "DigitalComm",
	2: "Common_V2",
}

// 0:CH  1:EN 2:TCH 3:JA 4:ES
var LangTypeMap = map[int]string{
	0: "CH",
	1: "EN",
	2: "TCH",
	3: "JA",
	4: "ES",
	5: "TH",
}

type WebsocketTokenReq struct {
	Token string `form:"token" json:"token" binding:"required"`
}
