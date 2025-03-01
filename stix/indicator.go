package stix

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator"
)

type Indicator struct {
	SDO
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	IndicatorTypes  []string  `json:"indicator_types"`
	Pattern         string    `json:"pattern"`
	PatternType     string    `json:"pattern_type"`
	PatternVersion  string    `json:"pattern_version"`
	ValidFrom       time.Time `json:"valid_from" `
	ValidUntil      time.Time `json:"valid_until"`
	KillChainPhases []string  `json:"kill_chain_phases"`
}

func (indicator *Indicator) Validate() error {
	validator := validator.New()
	return validator.Struct(indicator)
}

// FromJSONBytes reads JSON from a byte slice into the struct
func (indicator *Indicator) FromJSONBytes(data []byte) error {
	if err := json.Unmarshal(data, indicator); err != nil {
		return err
	}
	return indicator.Validate()
}

// ToJSONBytes converts the struct to a JSON byte slice
func (d *Indicator) ToJSONBytes() ([]byte, error) {
	return json.Marshal(d)
}
