package models

import "time"

type Trip struct {
	ID            uint64    `json:"id"`
	AccountID     uint64    `json:"accid"`
	Source        string    `json:"source"`
	Destination   string    `json:"dest"`
	Date          time.Time `json:"date"`
	Vehicle       string    `json:"vehicle"`
	Accomodation  uint64    `json:"accomodation"`
	CostPerPerson string    `json:"costperperson"`
}
