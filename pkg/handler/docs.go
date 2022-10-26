package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getDocs(c *gin.Context) {
	c.HTML(http.StatusOK, "docs.html", gin.H{
		"title": "Docs",
	})
}
