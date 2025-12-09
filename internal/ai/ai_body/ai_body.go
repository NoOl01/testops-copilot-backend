package ai_body

type AiBody struct {
	Model       string      `json:"model"`
	Messages    []AiMessage `json:"messages"`
	Temperature float64     `json:"temperature"`
	TopP        float64     `json:"top_p"`
	MaxTokens   int         `json:"max_tokens"`
}

type AiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CloudAnswer struct {
	Created int64 `json:"created"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
			Refusal string `json:"refusal"`
		} `json:"message"`
		StopReason string `json:"stop_reason"`
	} `json:"choices"`
}
