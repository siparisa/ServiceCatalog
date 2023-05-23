package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// BadRequest sends a Bad Request response with the provided error message.
func BadRequest(ctx *gin.Context, message string, details string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"details": details,
		"error":   message,
	})
}

// InternalServerError sends an Internal Server Error response with the provided error message.
func InternalServerError(ctx *gin.Context, message string, details string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error":   message,
		"details": details,
	})
}

// BadRequestWithDetails sends a Bad Request response with the provided error message and details.
func BadRequestWithDetails(ctx *gin.Context, message, details string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error":   message,
		"details": details,
	})
}

// OK sends a success response with the provided data.
func OK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

// NotFound sends a Not Found response with the provided error message.
func NotFound(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": message,
	})
}

// BuildErrorResponse builds an error response with the provided status code, error message, and details.
func BuildErrorResponse(ctx *gin.Context, message, details string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error":   message,
		"details": details,
	})
}
