package stix

import "time"

type SDO struct {
	Type               string    `json:"type" validate:"required"`
	SpecVersion        string    `json:"spec_version" validate:"required"`
	ID                 string    `json:"id" validate:"required"`
	CreatedByRef       bool      `json:"created_by_ref"`
	Created            time.Time `json:"created" `
	Modified           time.Time `json:"modified"`
	Revoked            bool      `json:"revoked"`
	Labels             []string  `json:"labels"`
	Confidence         int       `json:"confidence"`
	Lang               string    `json:"lang"`
	ExternalReferences []string  `json:"external_references"`
	ObjectMarkingRefs  []string  `json:"object_marking_refs"`
	GranularMarkings   []string  `json:"granular_markings"`
}
