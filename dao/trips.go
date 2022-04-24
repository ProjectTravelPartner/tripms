package dao

import (
	"database/sql"
	"fmt"

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

func TripsGetByAccID(accid uint64) ([]models.Trip, error) {
	qry := `
			SELECT 
				*
			FROM
				trips
			WHERE
				accid=?`
	fmt.Println(accid)
	rows, _ := dbclient.Query(qry, accid)
	var data []models.Trip
	defer rows.Close()
	for rows.Next() {
		o := getTripsFromRow(rows)
		data = append(data, o)
	}
	return data, nil
}

func TripsGetBySrcDest(source string, dest string) ([]models.Trip, error) {
	qry := `
			SELECT 
				*
			FROM
				trips
			WHERE
				source=?
			AND
				destination=?`
	fmt.Println(source)
	fmt.Println(dest)
	rows, _ := dbclient.Query(qry, source, dest)
	var data []models.Trip
	defer rows.Close()
	for rows.Next() {
		o := getTripsFromRow(rows)
		data = append(data, o)
	}
	return data, nil
}

func getTripsFromRow(row *sql.Rows) models.Trip {
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
	return data
}
