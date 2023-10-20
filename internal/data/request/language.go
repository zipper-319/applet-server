package request

type LabelInfo struct {
	Video          string `json:"video" binding:"required"`
	IsCorrect      bool   `json:"is_correct"`
	UltimateResult string `json:"universal_result"`
}
