package main

//go:generate ./bin/generator $GOFILE sdo
type Grouping struct {
	SDO
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Context     []string `json:"context" validate:"required"`
	ObjectRefs  []string `json:"object_refs" validate:"required"`
}

func (grouping *Grouping) GenerateID() string {
	return "grouping--" + v5UUID(grouping.Name+grouping.Description)
}
