package stix

import (
	"time"
)

//go:generate ../bin/generator $GOFILE sdo
type Indicator struct {
	SDO             `validate:"dive"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description" validate:"required"`
	Types           []string  `json:"indicator_types" validate:"required"`
	Pattern         string    `json:"pattern" validate:"required"`
	PatternType     string    `json:"pattern_type" validate:"required,oneof=stix snort yara"`
	PatternVersion  string    `json:"pattern_version"`
	ValidFrom       time.Time `json:"valid_from" validate:"required"`
	ValidUntil      time.Time `json:"valid_until"`
	KillChainPhases []string  `json:"kill_chain_phases"`
}

func (indicator *Indicator) GenerateID() string {
	return "indicator--" + v5UUID(indicator.Name+indicator.Description+indicator.Pattern+indicator.PatternType)
}
