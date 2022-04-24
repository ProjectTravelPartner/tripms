package api

import (
	"fmt"
	"strconv"

	"github.com/ProjectTravelPartner/commonclient"
	"github.com/ProjectTravelPartner/tripms/dao"
	"github.com/ProjectTravelPartner/tripms/models"
	"github.com/gin-gonic/gin"
)

func TripCreate(c *gin.Context) {
	var data models.Trip
	c.BindJSON(&data)
	fmt.Println("#%v", data)
	id, _ := dao.TripCreate(data)
	commonclient.RespOK(c, id)
}

func TripGet(c *gin.Context) {
	var out models.Trip
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	out, _ = dao.TripGet(id)
	commonclient.RespOK(c, out)
}
