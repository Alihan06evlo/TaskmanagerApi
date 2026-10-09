package task

import (

	"net/http"
	"strconv"
	

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Delete(c *gin.Context) {
 idParam := c.Param("id")
 id, err := strconv.Atoi(idParam)
 if err != nil {
  logrus.WithError(err).Warn("[Validation] your id is not int")
  c.Status(http.StatusBadRequest)
  return
 }


 err = h.srv.Delete(id)
 if err != nil {
  logrus.WithError(err).Warn("[Delete] not found")
  c.Status(http.StatusInternalServerError) 
  return
 }

 c.Status(http.StatusNoContent)
}