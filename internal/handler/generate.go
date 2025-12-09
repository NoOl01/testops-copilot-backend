package handler

import (
	"context"
	"net/http"
	"testops_copilot/internal/consts"
	"testops_copilot/internal/dto"
	"testops_copilot/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// Generate
// @Summary Generate test case using LLM
// @Description Generates a test case based on the user prompt. The request will be processed by the LLM and returned in a unified response format.
// @Tags Generate
// @Accept json
// @Produce json
// @Param body body dto.GenerateBody true "Request body containing test cases"
// @Success 200 {object} dto.GenerateResult "Successful response with generated test case"
// @Failure 400 {object} dto.ErrorResult "Invalid request body"
// @Failure 500 {object} dto.ErrorResult "Internal server error or LLM request failed"
// @Router /api/v1/generate [post]
func (h *handler) Generate(ctx *gin.Context) {
	timeoutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 120*time.Second)
	defer cancel()

	logger.Log.Debug(consts.GenerateHandler, "checking the request body")
	var body dto.GenerateBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		logger.Log.Debug(consts.GenerateHandler, "invalid request body")

		ctx.JSON(http.StatusBadRequest, dto.ErrorResult{
			Status: consts.BadRequest,
			Error:  consts.GenerateInvalidBody.Error(),
		})
		return
	}

	logger.Log.Debug(consts.GenerateHandler, "body is valid")

	//wg := sync.WaitGroup{}
	//
	//wg.Add(len(body.Cases))
	//for _, testCase := range body.Cases {
	//	go
	//}

	var response *dto.GenerateResult
	response, err := h.Service.Generate(body.Cases[0], timeoutCtx)
	if err != nil {
		ctx.JSON(500, dto.ErrorResult{
			Status: "error",
			Error:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, *response)
}
