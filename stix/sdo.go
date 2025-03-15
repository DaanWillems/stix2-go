package stix

import "time"

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

// Common SDO option functions
type SDOOption func(*SDO)

func WithID(id string) SDOOption {
	return func(s *SDO) {
		s.ID = id
	}
}

func WithSpecVersion(version string) SDOOption {
	return func(s *SDO) {
		s.SpecVersion = version
	}
}

func WithCreatedByRef(ref bool) SDOOption {
	return func(s *SDO) {
		s.CreatedByRef = ref
	}
}

func WithCreated(t time.Time) SDOOption {
	return func(s *SDO) {
		s.Created = t
	}
}

func WithModified(t time.Time) SDOOption {
	return func(s *SDO) {
		s.Modified = t
	}
}

func WithRevoked(revoked bool) SDOOption {
	return func(s *SDO) {
		s.Revoked = revoked
	}
}

func WithLabels(labels []string) SDOOption {
	return func(s *SDO) {
		s.Labels = labels
	}
}

func WithConfidence(confidence int) SDOOption {
	return func(s *SDO) {
		s.Confidence = confidence
	}
}

func WithLang(lang string) SDOOption {
	return func(s *SDO) {
		s.Lang = lang
	}
}

func WithExternalReferences(refs []ExternalReference) SDOOption {
	return func(s *SDO) {
		s.ExternalReferences = refs
	}
}

func WithObjectMarkingRefs(refs []string) SDOOption {
	return func(s *SDO) {
		s.ObjectMarkingRefs = refs
	}
}

func WithGranularMarkings(markings []string) SDOOption {
	return func(s *SDO) {
		s.GranularMarkings = markings
	}
}

func applySDOOptions(sdo *SDO, options []SDOOption) {
	for _, option := range options {
		option(sdo)
	}
}

func CommonSDOOptions(options ...SDOOption) []SDOOption {
	return options
}
