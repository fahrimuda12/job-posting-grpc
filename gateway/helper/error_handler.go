package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SwitchStatusToCode(status int) int {
	switch status {
	case 0:
		return 200
	case 1:
		return 403
	case 2:
		return 520
	case 3:
		return 400 
	case 4:
		return 500
	case 5:
		return 404
	default:
		return 500
	}
}

func SuccessResponse(data interface{}, message string) gin.H {
	return gin.H{"status": http.StatusOK, "data": data, "message":message,"error": false}
}

func ErrorResponse(message string, code string) gin.H {
	return gin.H{"error": true, "message": message, "status": code}
}

func ServerErrorResponse(error string) gin.H {
	return gin.H{"error": error, "message": "Server Error", "status": http.StatusInternalServerError}
}