package stix

import (
	"time"
)

type ThreatActor struct {
	SDO                 `json:"sdo,omitempty"`
	Name                string    `json:"name,omitempty"`
	Description         string    `json:"description,omitempty"`
	Aliases             []string  `json:"aliases,omitempty"`
	FirstSeen           time.Time `json:"first_seen,omitempty"`
	LastSeen            time.Time `json:"last_seen,omitempty"`
	Roles               []string  `json:"roles,omitempty"`
	Goals               []string  `json:"goals,omitempty"`
	Sophistication      string    `json:"sophistication,omitempty"`
	ResourceLevel       string    `json:"resource_level,omitempty"`
	PrimaryMotivation   string    `json:"primary_motivation,omitempty"`
	SecondaryMotivation []string  `json:"secondary_motivation,omitempty"`
	PersonalMotivation  []string  `json:"personal_motivation,omitempty"`
}
