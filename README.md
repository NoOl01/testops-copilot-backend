# TestOps Copilot

TestOps Copilot is an internal service for generating test cases using large language models (LLMs).
Built in Go with Gin, it provides a REST API, handles long-running requests safely, and integrates with cloud LLMs (e.g., Cloud.ru).

## Tech Stack

Backend: Go 1.24, Gin
LLM Integration: Cloud.ru Foundation Models
Logging: Custom logger `golog`
Deployment: Docker, env configuration

## Getting Started

### Required

- Go 1.24 +
- Docker

### Building

- Get API key from [Cloud.ru](https://cloud.ru/docs/console_api/ug/topics/guides__static-api-keys)
- Clone repository `git clone https://github.com/NoOl01/testops-copilot-backend`
- Create .env file in the root directory:
  ```LLM_TOKEN=your_api_key_here
  MODEL=Qwen/Qwen3-Coder-480B-Instruct
  LLM_URL=https://foundation-models.api.cloud.ru/v1/chat/completions
  SERVER_PORT=8080
  DEBUG=true
  TEMPERATURE=0.5
  TOP_P=0.9
  MAX_TOKENS=5000
- go to the root directory (cmd/powershell)
- run build: `docker-compose up --build`
