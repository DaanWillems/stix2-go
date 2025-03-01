package stix

import "time"

type Campaign struct {
	SDO
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Aliases     []string  `json:"aliases"`
	FirstSeen   time.Time `json:"first_seen"`
	LastSeen    time.Time `json:"last_seen"`
	Objective   string    `json:"objective"`
}
