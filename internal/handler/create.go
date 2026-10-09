package task

import (
	"net/http"
	"taskmanagerApi/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Create(c *gin.Context) {
	var input *entity.TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.WithError(err).Warn("[Validator] invalid json")
		c.Status(http.StatusBadRequest)
		return
	}
	task, err := h.srv.Create(c.Request.Context(),input)
	if err != nil {
		logrus.WithError(err).Warn("[Create] something went wrong")
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusCreated, task)
}