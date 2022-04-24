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

func TripsGet(c *gin.Context) {
	var out []models.Trip
	source := c.Query("src")
	dest := c.Query("dest")
	if source != "" && dest != "" {
		out, _ = dao.TripsGetBySrcDest(source, dest)
	}
	accid := c.Query("accid")
	if accid != "" {
		accidint, _ := strconv.ParseUint(accid, 10, 64)
		out, _ = dao.TripsGetByAccID(accidint)
	}
	commonclient.RespOK(c, out)
}
