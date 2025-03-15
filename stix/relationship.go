package stix

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator"
)

type Relationship struct {
	SRO              `validate:"dive"`
	RelationshipType string    `json:"relationship_type" validate:"required"` //TODO: Enforce type
	Description      string    `json:"description"`
	SourceRef        string    `json:"source_ref" validate:"required"`
	TargetRef        string    `json:"target_ref" validate:"required"`
	StartTime        time.Time `json:"start_time"`
	StopTime         time.Time `json:"stop_time"`
}

func (relationship *Relationship) GenerateID() string {
	return "relationship--" + v5UUID(relationship.SourceRef+relationship.Description+relationship.TargetRef)
}

func (Relationship *Relationship) Validate() error {
	validator := validator.New()
	return validator.Struct(Relationship)
}

// FromJSONBytes reads JSON from a byte slice into the struct
func (Relationship *Relationship) FromJSONBytes(data []byte) error {
	if err := json.Unmarshal(data, Relationship); err != nil {
		return err
	}
	return Relationship.Validate()
}

// ToJSONBytes converts the struct to a JSON byte slice
func (r *Relationship) ToJSONBytes() ([]byte, error) {
	return json.Marshal(r)
}

type RelationshipOption func(*Relationship)

func WithRelationshipType(relationshipType string) RelationshipOption {
	return func(i *Relationship) {
		i.RelationshipType = relationshipType
	}
}

func WithRelationshipDescription(description string) RelationshipOption {
	return func(i *Relationship) {
		i.Description = description
	}
}

func WithTargetRef(targetRef string) RelationshipOption {
	return func(i *Relationship) {
		i.TargetRef = targetRef
	}
}

func WithSourceRef(sourceRef string) RelationshipOption {
	return func(i *Relationship) {
		i.SourceRef = sourceRef
	}
}

func WithStartTime(startTime time.Time) RelationshipOption {
	return func(i *Relationship) {
		i.StartTime = startTime
	}
}

func WithStopTime(stopTime time.Time) RelationshipOption {
	return func(i *Relationship) {
		i.StopTime = stopTime
	}
}

func NewRelationship(sroOptions []SROOption, RelationshipOptions []RelationshipOption) *Relationship {
	now := time.Now()

	// Create base SDO with required fields
	sro := SRO{
		Type:        "Relationship",
		SpecVersion: "2.1", // Default version
		Created:     now,
		Modified:    now,
	}

	// Apply SDO options
	applySROOptions(&sro, sroOptions)

	// Create Relationship with embedded SDO
	relationship := &Relationship{
		SRO: sro,
	}

	// Apply Relationship-specific options
	for _, option := range RelationshipOptions {
		option(relationship)
	}

	return relationship
}
