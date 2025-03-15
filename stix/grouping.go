package stix

import (
	"time"

	"github.com/go-playground/validator"
)

type Grouping struct {
	SDO
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Context     []string `json:"context" validate:"required"`
	ObjectRefs  []string `json:"object_refs" validate:"required"`
}

func (grouping *Grouping) Validate() error {
	validator := validator.New()
	return validator.Struct(grouping)
}

func (grouping *Grouping) GenerateID() string {
	return "grouping--" + v5UUID(grouping.Name+grouping.Description)
}

type GroupingOption func(*Grouping)

func WithGroupingName(name string) GroupingOption {
	return func(g *Grouping) {
		g.Name = name
	}
}

func WithGroupingDescription(description string) GroupingOption {
	return func(g *Grouping) {
		g.Description = description
	}
}

func WithContext(context []string) GroupingOption {
	return func(g *Grouping) {
		g.Context = context
	}
}

func WithObjectRefs(objectRefs []string) GroupingOption {
	return func(g *Grouping) {
		g.ObjectRefs = objectRefs
	}
}

func NewGrouping(sdoOptions []SDOOption, groupingOptions []GroupingOption) *Grouping {
	now := time.Now()

	// Create base SDO with required fields
	sdo := SDO{
		Type:        "groupinh",
		SpecVersion: "2.1", // Default version
		Created:     now,
		Modified:    now,
	}

	for _, option := range sdoOptions {
		option(&sdo)
	}

	grouping := &Grouping{
		SDO: sdo,
	}

	for _, option := range groupingOptions {
		option(grouping)
	}

	return grouping
}
