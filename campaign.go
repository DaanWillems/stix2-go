package stix

import (
	"time"
)

//go:generate ./bin/generator $GOFILE sdo
type Campaign struct {
	SDO
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Aliases     []string  `json:"aliases"`
	FirstSeen   time.Time `json:"first_seen"`
	LastSeen    time.Time `json:"last_seen"`
	Objective   string    `json:"objective"`
}

func (campaign *Campaign) GenerateID() string {
	return "campaign--" + v5UUID(campaign.Name+campaign.Description)
}
