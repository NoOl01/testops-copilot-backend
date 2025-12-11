package prompts

import (
	"bytes"
	_ "embed"
	"testops_copilot/internal/consts"
	"testops_copilot/internal/dto"
	"testops_copilot/pkg/logger"
	"text/template"
)

//go:embed templates/ui_test.tmpl
var uiSystemPrompt string

//go:embed templates/api_test.tmpl
var apiSystemPrompt string

func BuildPrompt(testCase dto.Case) (string, error) {
	logger.Log.Debug(consts.PromptBuilder, "Building prompt")
	var tmplSource string

	switch testCase.TestType {
	case dto.UiTest:
		tmplSource = uiSystemPrompt
	case dto.ApiTest:
		tmplSource = apiSystemPrompt
	default:
		return "", consts.InvalidTestType
	}

	tmpl, err := template.New("systemPrompt").Parse(tmplSource)
	if err != nil {
		logger.Log.Error(consts.PromptBuilder, "error with parsing template: "+err.Error())
		return "", err
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, testCase); err != nil {
		logger.Log.Error(consts.PromptBuilder, "building template failed: "+err.Error())
		return "", err
	}

	return out.String(), nil
}
