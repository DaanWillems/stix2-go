package main

import "time"

//go:generate ./bin/generator $GOFILE
type SDO struct {
	Type               string              `json:"type" validate:"required"`
	SpecVersion        string              `json:"spec_version" validate:"required"`
	ID                 string              `json:"id" validate:"required"`
	CreatedByRef       bool                `json:"created_by_ref"`
	Created            time.Time           `json:"created" validate:"required"`
	Modified           time.Time           `json:"modified" validate:"required"`
	Revoked            bool                `json:"revoked"`
	Labels             []string            `json:"labels"`
	Confidence         int                 `json:"confidence"`
	Lang               string              `json:"lang"`
	ExternalReferences []ExternalReference `json:"external_references"`
	ObjectMarkingRefs  []string            `json:"object_marking_refs"`
	GranularMarkings   []string            `json:"granular_markings"`
}

func applySDOOptions(sdo *SDO, options []SDOOption) {
	for _, option := range options {
		option(sdo)
	}
}

func CommonSDOOptions(options ...SDOOption) []SDOOption {
	return options
}
