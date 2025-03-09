package stix

import (
	"time"

	"github.com/go-playground/validator"
)

type ThreatActor struct {
	SDO                 `json:"sdo,omitempty"`
	Name                string    `json:"name" validate:"required"`
	Description         string    `json:"description,omitempty"`
	ThreatActorTypes    []string  `json:"threat_actor_types" validate:"required"`
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

func (threatActor *ThreatActor) Validate() error {
	validator := validator.New()
	return validator.Struct(threatActor)
}

// ThreatActorOption defines a functional option for configuring a ThreatActor
type ThreatActorOption func(*ThreatActor)

func WithThreatActorName(name string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.Name = name
	}
}

func WithThreatActorDescription(description string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.Description = description
	}
}
func WithThreatActorTypes(threatActorTypes []string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.ThreatActorTypes = threatActorTypes
	}
}

func WithAliases(aliases []string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.Aliases = aliases
	}
}

func WithFirstSeen(t time.Time) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.FirstSeen = t
	}
}

func WithLastSeen(t time.Time) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.LastSeen = t
	}
}

func WithRoles(roles []string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.Roles = roles
	}
}

func WithGoals(goals []string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.Goals = goals
	}
}

func WithSophistication(sophistication string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.Sophistication = sophistication
	}
}

func WithResourceLevel(resourceLevel string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.ResourceLevel = resourceLevel
	}
}

func WithPrimaryMotivation(motivation string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.PrimaryMotivation = motivation
	}
}

func WithSecondaryMotivation(motivations []string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.SecondaryMotivation = motivations
	}
}

func WithPersonalMotivation(motivations []string) ThreatActorOption {
	return func(ta *ThreatActor) {
		ta.PersonalMotivation = motivations
	}
}

// NewThreatActor creates a new ThreatActor with the given ID and options
func NewThreatActor(id string, sdoOptions []SDOOption, threatActorOptions []ThreatActorOption) *ThreatActor {
	now := time.Now()

	// Create base SDO with required fields
	sdo := SDO{
		Type:        "threat-actor",
		ID:          id,
		SpecVersion: "2.1", // Default version
		Created:     now,
		Modified:    now,
	}

	// Apply SDO options
	for _, option := range sdoOptions {
		option(&sdo)
	}

	// Create threat actor with embedded SDO
	threatActor := &ThreatActor{
		SDO: sdo,
	}

	// Apply threat actor-specific options
	for _, option := range threatActorOptions {
		option(threatActor)
	}

	return threatActor
}
