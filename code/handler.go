package code

import (
	"context"
	"golox-playground/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Run(ctx context.Context, code Code) (Output, error)
}

type handler struct {
	service Service
}

func NewHandler(s Service) handler {
	return handler{
		service: s,
	}
}

func (h *handler) Run(ctx *gin.Context) {
	var code Code

	err := ctx.ShouldBind(&code)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{
				Message: "error binding code",
				Data:    nil,
			},
		)
		return
	}

	output, err := h.service.Run(ctx.Request.Context(), code)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.Response{
				Message: "error binding code",
				Data:    nil,
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		response.Response{
			Message: "code ran successfully",
			Data:    output,
		},
	)
}
