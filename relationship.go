package main

import (
	"time"
)

//go:generate ./bin/generator $GOFILE sro
type Relationship struct {
	SRO         `validate:"dive"`
	Type        string    `json:"relationship_type" validate:"required"` //TODO: Enforce type
	Description string    `json:"description"`
	SourceRef   string    `json:"source_ref" validate:"required"`
	TargetRef   string    `json:"target_ref" validate:"required"`
	StartTime   time.Time `json:"start_time"`
	StopTime    time.Time `json:"stop_time"`
}

func (relationship *Relationship) GenerateID() string {
	return "relationship--" + v5UUID(relationship.SourceRef+relationship.Description+relationship.TargetRef)
}
