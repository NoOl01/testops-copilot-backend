package dto

type ErrorResult struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type GenerateResult struct {
	Created    int64   `json:"created"`
	Content    string  `json:"content"`
	Refusal    *string `json:"refusal"`
	StopReason string  `json:"stop_reason"`
}
