package stix

import (
	"time"
)

//go:generate ./bin/generator $GOFILE sdo
type Infrastructure struct {
	SDO             `validate:"dive"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description"`
	Types           []string  `json:"infrastructure_types"`
	Aliases         string    `json:"aliases"`
	KillChainPhases []string  `json:"kill_chain_phases"`
	FirstSeen       time.Time `json:"first_seen"`
	LastSeen        time.Time `json:"last_seen"`
}

func (infrastructure *Infrastructure) GenerateID() string {
	return "indicator--" + v5UUID(infrastructure.Name+infrastructure.Description)
}
