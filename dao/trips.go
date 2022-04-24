package dao

import (
	"github.com/ProjectTravelPartner/dbclient"
	"github.com/ProjectTravelPartner/tripms/models"
)

func TripCreate(data models.Trip) (uint64, error) {
	qry := `
			INSERT INTO trips
				(accid,source,destination,travel_date,vehicle,accomodation,costperperson)
			VALUES
				(?,?,?,?,?,?,?)`
	id, err := dbclient.ExecGetID(qry, data.AccountID, data.Source, data.Destination, data.Date, data.Vehicle, data.Accomodation, data.CostPerPerson)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func TripGet(id uint64) (models.Trip, error) {
	qry := `
			SELECT 
				*
			FROM
				trips
			WHERE
				id=?`
	row := dbclient.QueryRow(qry, id)
	var data models.Trip
	row.Scan(
		&data.ID,
		&data.AccountID,
		&data.Source,
		&data.Destination,
		&data.Date,
		&data.Vehicle,
		&data.Accomodation,
		&data.CostPerPerson,
	)
	return data, nil

}
