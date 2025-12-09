package utils

import (
	"testops_copilot/internal/ai/ai_body"
	"testops_copilot/internal/dto"
)

func AnswerToGenerateResult(ans ai_body.CloudAnswer) dto.GenerateResult {
	var content string
	var stopReason string
	var refusal *string

	if len(ans.Choices) > 0 {
		content = ans.Choices[0].Message.Content
		stopReason = ans.Choices[0].StopReason
		refusal = Str(ans.Choices[0].Message.Refusal)
	}

	return dto.GenerateResult{
		Created:    ans.Created,
		Content:    content,
		Refusal:    refusal,
		StopReason: stopReason,
	}
}
