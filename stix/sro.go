package stix

import "time"

type SRO struct {
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

// Common SRO option functions
type SROOption func(*SRO)

func WithSROSpecVersion(version string) SROOption {
	return func(s *SRO) {
		s.SpecVersion = version
	}
}

func WithSROCreatedByRef(ref bool) SROOption {
	return func(s *SRO) {
		s.CreatedByRef = ref
	}
}

func WithSROCreated(t time.Time) SROOption {
	return func(s *SRO) {
		s.Created = t
	}
}

func WithSROModified(t time.Time) SROOption {
	return func(s *SRO) {
		s.Modified = t
	}
}

func WithSRORevoked(revoked bool) SROOption {
	return func(s *SRO) {
		s.Revoked = revoked
	}
}

func WithSROLabels(labels []string) SROOption {
	return func(s *SRO) {
		s.Labels = labels
	}
}

func WithSROConfidence(confidence int) SROOption {
	return func(s *SRO) {
		s.Confidence = confidence
	}
}

func WithSROLang(lang string) SROOption {
	return func(s *SRO) {
		s.Lang = lang
	}
}

func WithSROExternalReferences(refs []ExternalReference) SROOption {
	return func(s *SRO) {
		s.ExternalReferences = refs
	}
}

func WithSROObjectMarkingRefs(refs []string) SROOption {
	return func(s *SRO) {
		s.ObjectMarkingRefs = refs
	}
}

func WithSROGranularMarkings(markings []string) SROOption {
	return func(s *SRO) {
		s.GranularMarkings = markings
	}
}

func applySROOptions(SRO *SRO, options []SROOption) {
	for _, option := range options {
		option(SRO)
	}
}

func CommonSROOptions(options ...SROOption) []SROOption {
	return options
}
