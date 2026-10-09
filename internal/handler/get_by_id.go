package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

 func (h *Handler) GetById(c *gin.Context) {
	idParam := c.Param("id")
	id,err := strconv.Atoi(idParam) 
	if err != nil {
		logrus.WithError(err).Error("[Service] your id is not int")
		c.Status(http.StatusBadRequest)
		return
	} 
	task,err := h.srv.GetById(id)
	if err != nil {
		logrus.WithError(err).Error("[Service] task not found ")
		c.Status(http.StatusNotFound)
		return 
	}
	c.JSON(http.StatusOK, task)
 } 
