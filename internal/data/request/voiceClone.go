package request

type VoiceUploadReq struct {
	Sequence int `form:"sequence" json:"sequence,omitempty"`
}

type VoiceUploadRes struct {
	NextSequence int `form:"nextSequence" json:"nextSequence,omitempty"`
}

type ProgressRes struct {
	CurrentNumber int    `form:"currentNumber" json:"currentNumber"`
	FinishedTime  string `form:"finishedTime" json:"finishedTime"`
	AwaitTrain    int    `form:"awaitTrain" json:"awaitTrain,omitempty"`
}

type DownloadReq struct {
	Sequence int `form:"sequence" json:"sequence,omitempty"`
}

type CommitRes struct {
	FinishedTime string `form:"finishedTime" json:"finishedTime,omitempty"`
	AwaitTrain   int    `form:"awaitTrain" json:"awaitTrain,omitempty"`
}
