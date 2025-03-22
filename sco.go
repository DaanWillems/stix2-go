package stix

//go:generate ./bin/generator $GOFILE
type SCO struct {
	Type              string   `json:"type" validate:"required"`
	SpecVersion       string   `json:"spec_version,omitempty"`
	ID                string   `json:"id" validate:"required"`
	ObjectMarkingRefs []string `json:"object_marking_refs,omitempty"`
	GranularMarkings  []string `json:"granular_markings,omitempty"`
}

func applySCOOptions(SCO *SCO, options []SCOOption) {
	for _, option := range options {
		option(SCO)
	}
}

func CommonSCOOptions(options ...SCOOption) []SCOOption {
	return options
}
