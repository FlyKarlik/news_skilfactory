package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetNews(c *gin.Context) {
	count := c.Param("count")
	if count == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Empty count",
		})
		return
	}

	countInt, err := strconv.Atoi(count)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := h.svc.NewsService.GetNews(countInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": response,
	})
}
