package stix

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator"
)

type Indicator struct {
	SDO             `validate:"dive"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description" validate:"required"`
	IndicatorTypes  []string  `json:"indicator_types" validate:"required"`
	Pattern         string    `json:"pattern" validate:"required"`
	PatternType     string    `json:"pattern_type" validate:"required,oneof=stix snort yara"`
	PatternVersion  string    `json:"pattern_version"`
	ValidFrom       time.Time `json:"valid_from" validate:"required"`
	ValidUntil      time.Time `json:"valid_until"`
	KillChainPhases []string  `json:"kill_chain_phases"`
}

func (indicator *Indicator) Validate() error {
	validator := validator.New()
	return validator.Struct(indicator)
}

func (indicator *Indicator) GenerateID() string {
	return "indicator--" + v5UUID(indicator.Name+indicator.Description+indicator.Pattern+indicator.PatternType)
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

type IndicatorOption func(*Indicator)

func WithIndicatorName(name string) IndicatorOption {
	return func(i *Indicator) {
		i.Name = name
	}
}

func WithIndicatorDescription(description string) IndicatorOption {
	return func(i *Indicator) {
		i.Description = description
	}
}

func WithIndicatorTypes(types []string) IndicatorOption {
	return func(i *Indicator) {
		i.IndicatorTypes = types
	}
}

func WithPattern(pattern string) IndicatorOption {
	return func(i *Indicator) {
		i.Pattern = pattern
	}
}

func WithPatternType(patternType string) IndicatorOption {
	return func(i *Indicator) {
		i.PatternType = patternType
	}
}

func WithPatternVersion(version string) IndicatorOption {
	return func(i *Indicator) {
		i.PatternVersion = version
	}
}

func WithValidFrom(t time.Time) IndicatorOption {
	return func(i *Indicator) {
		i.ValidFrom = t
	}
}

func WithValidUntil(t time.Time) IndicatorOption {
	return func(i *Indicator) {
		i.ValidUntil = t
	}
}

func WithKillChainPhases(phases []string) IndicatorOption {
	return func(i *Indicator) {
		i.KillChainPhases = phases
	}
}

func NewIndicator(sdoOptions []SDOOption, indicatorOptions []IndicatorOption) *Indicator {
	now := time.Now()

	// Create base SDO with required fields
	sdo := SDO{
		Type:        "indicator",
		SpecVersion: "2.1", // Default version
		Created:     now,
		Modified:    now,
	}

	// Apply SDO options
	applySDOOptions(&sdo, sdoOptions)

	// Create indicator with embedded SDO
	indicator := &Indicator{
		SDO: sdo,
	}

	// Apply indicator-specific options
	for _, option := range indicatorOptions {
		option(indicator)
	}

	return indicator
}
