package main

import (
	"fmt"

	"github.com/ProjectTravelPartner/commonclient"
	"github.com/ProjectTravelPartner/dbclient"
	"github.com/ProjectTravelPartner/tripms/api"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hey, there!")
	port := commonclient.CliPort(3000)
	dbclient.Init()
	defer dbclient.Close()
	router := gin.New()
	group := router.Group("tripmgts")

	group.GET("/trip/:id", api.TripGet)
	group.GET("/trips", api.TripsGet)
	group.POST("/trip", api.TripCreate)
	group.PUT("/trip", api.TripUpdate)
	router.Run(fmt.Sprintf(":%v", port))
	//group.DELETE("/account/:id", api.AccountDelete)
}
