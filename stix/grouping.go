package stix

type Grouping struct {
	SDO
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Context     []string `json:"context "`
	ObjectRefs  []string `json:"object_refs"`
}
