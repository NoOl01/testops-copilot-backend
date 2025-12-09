package dto

type GenerateBody struct {
	Cases []Case `json:"cases"`
}

type Case struct {
	UserPrompt     string         `json:"user_prompt"`
	AllureMetadata AllureMetadata `json:"allure_metadata"`
}

type AllureMetadata struct {
	Manual  bool     `json:"manual"`
	Feature string   `json:"feature"`
	Story   string   `json:"story"`
	Label   string   `json:"label"`
	Tag     []string `json:"tag"`
	Title   string   `json:"title"`
}
