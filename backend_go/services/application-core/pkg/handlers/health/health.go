package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHealth(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
