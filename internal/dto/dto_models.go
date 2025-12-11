package dto

type GenerateBody struct {
	Cases []Case `json:"cases"`
}

type Case struct {
	UserPrompt     string         `json:"user_prompt"`
	TestType       string         `json:"test_type"`
	AllureMetadata AllureMetadata `json:"allure_metadata"`
}

type AllureMetadata struct {
	Manual   bool   `json:"manual"`
	Label    string `json:"label"`
	Feature  string `json:"feature"`
	Story    string `json:"story"`
	Title    string `json:"title"`
	Link     string `json:"link"`
	Tags     string `json:"tags"`
	Priority string `json:"priority"`
}

const (
	UiTest  string = "UiTest"
	ApiTest string = "ApiTest"
)
