package request

type TTSCallReqData struct {
	// 发音人
	Speaker string `form:"speaker" json:"speaker"  binding:"required"`
	// 速度
	Speed int `form:"speed" json:"speed"`
	// 音量
	Volume int `form:"volume" json:"volume"`
	// 1:low   2:medium  3:high  默认2:medium
	Pitch int `form:"pitch" json:"pitch"`
	// 文本
	Text string `form:"text" json:"text,omitempty"  binding:"required"`
}

type TTSReqParam struct {
	// cloud
	Location string `form:"location" json:"location"`
	// cloudminds
	Type string `form:"type" json:"type"`
	// 发音人
	Speaker int `form:"speaker" json:"speaker"  binding:"required"`
	// 速度
	Speed int `form:"speed" json:"speed"`
	// 音量
	Volume int `form:"volume" json:"volume"`
	// 1:low   2:medium  3:high  默认2:medium
	Pitch int `form:"pitch" json:"pitch"`
	// 是否流式
	StreamEnable bool `form:"streamEnable" json:"streamEnable"  binding:"required"`
	// 是否需要预处理
	TextPreHandle bool `form:"textPreHandle" json:"textPreHandle"  binding:"required"`
	// voiceTuning开关, true:on/ false: off
	VoiceTuning bool `form:"voiceTuning" json:"voiceTuning"   binding:"required"`
}

var VoiceTuningMap = map[bool]string{
	true:  "on",
	false: "off",
}

var PitchMap = map[int]string{
	1: "low",
	2: "medium",
	3: "high",
}
