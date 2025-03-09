package stix

type SCO struct {
	Type              string   `json:"type" validate:"required"`
	SpecVersion       string   `json:"spec_version,omitempty"`
	ID                string   `json:"id" validate:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings  []string `json:"granular_markings,omitempty"`
}

// Common SCO option functions
type SCOOption func(*SCO)

func WithSCOSpecVersion(version string) SCOOption {
	return func(s *SCO) {
		s.SpecVersion = version
	}
}

func WithSCOObjectMarkingRefs(refs []string) SCOOption {
	return func(s *SCO) {
		s.ObjectMarkingRefs = refs
	}
}

func WithSCOGranularMarkings(markings []string) SCOOption {
	return func(s *SCO) {
		s.GranularMarkings = markings
	}
}

func applySCOOptions(SCO *SCO, options []SCOOption) {
	for _, option := range options {
		option(SCO)
	}
}

func CommonSCOOptions(options ...SCOOption) []SCOOption {
	return options
}
