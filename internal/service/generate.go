package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testops_copilot/internal/ai/ai_body"
	"testops_copilot/internal/ai/prompts"
	"testops_copilot/internal/config"
	"testops_copilot/internal/consts"
	"testops_copilot/internal/dto"
	"testops_copilot/internal/utils"
	"testops_copilot/pkg/logger"
)

func (s service) Generate(testCase dto.Case, ctx context.Context) (*dto.GenerateResult, error) {
	var systemPrompt string
	switch testCase.TestType {
	case dto.UiTest:
		systemPrompt = prompts.UiTest
	case dto.ApiTest:
		systemPrompt = prompts.ApiTest
	default:
		logger.Log.Error(consts.GenerateService, "invalid test type")
		return nil, consts.InvalidTestType
	}

	rawJson := ai_body.AiBody{
		Model: config.Env.Model,
		Messages: []ai_body.AiMessage{
			{
				Role:    "user",
				Content: testCase.UserPrompt,
			},
			{
				Role:    "system",
				Content: systemPrompt,
			},
		},
		Temperature: config.Env.Temperature,
		TopP:        config.Env.TopP,
		MaxTokens:   config.Env.MaxTokens,
	}

	jsonData, err := json.Marshal(rawJson)
	if err != nil {
		logger.Log.Error(consts.GenerateService, "error with generating json: "+err.Error())
		return nil, fmt.Errorf("error with generating json: %w", err)
	}

	jsonBody := bytes.NewBuffer(jsonData)

	logger.Log.Debug(consts.GenerateService, "building request")

	req, err := http.NewRequestWithContext(ctx, "POST", config.Env.LLMUrl, jsonBody)
	if err != nil {
		logger.Log.Info(consts.GenerateService, "request creation error: "+err.Error())
		return nil, fmt.Errorf("request creation error: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Env.LLMToken)

	logger.Log.Debug(consts.GenerateService, "sending request")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			logger.Log.Error(consts.GenerateService, "connection timeout")
			return nil, consts.ConnectionTimeout
		}
		if errors.Is(ctx.Err(), context.Canceled) {
			logger.Log.Error(consts.GenerateService, "request canceled")
			return nil, consts.RequestCanceled
		}
		logger.Log.Info(consts.GenerateService, "request error: "+err.Error())
		return nil, fmt.Errorf("request error: %w", err)
	}

	defer resp.Body.Close()

	logger.Log.Debug(consts.GenerateService, fmt.Sprintf("received response: %d", resp.StatusCode))

	var generateResp ai_body.CloudAnswer
	var result dto.GenerateResult
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Log.Error(consts.GenerateService, "error with reading error: "+err.Error())
			return nil, fmt.Errorf("error with reading error: %w", err)
		}
		logger.Log.Error(consts.GenerateService, fmt.Sprintf("AI API Error (status: %d): %s", resp.StatusCode, string(body)))
		return nil, fmt.Errorf("AI API Error (status: %d): %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(&generateResp); err != nil {
		logger.Log.Error(consts.GenerateService, "decode error: "+err.Error())
		return nil, fmt.Errorf("decode error %w", err)
	}

	result = utils.AnswerToGenerateResult(generateResp)

	return &result, nil
}
