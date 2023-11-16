package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowErrorMessageGin(c *gin.Context, message string) {
	c.HTML(
		http.StatusInternalServerError,
		"error.html",
		gin.H{
			"message": message,
			"title":   "Streams",
		},
	)
}

func ShowErrorGin(c *gin.Context, err error) {
	ShowErrorMessageGin(c, err.Error())
}

func RedirectGin(c *gin.Context, targetUrl string) {
	c.Redirect(http.StatusFound, targetUrl)
	c.Abort()
}
