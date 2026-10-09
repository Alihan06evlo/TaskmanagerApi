package task

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) List(c *gin.Context) {
	tasks, err := h.srv.GetAll()
	if err != nil {
		logrus.WithError(err).Error("[Service] something went wrong")
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, tasks)
}