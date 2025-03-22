package main

//go:generate ./bin/generator $GOFILE sdo
type Identity struct {
	SDO
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	Roles              []string `json:"roles"`
	IdentityClass      string   `json:"identity_class"`
	Sectors            []string `json:"sectors"`
	ContactInformation string   `json:"contact_information"`
}

func (identity *Identity) GenerateID() string {
	return "grouping--" + v5UUID(identity.Name+identity.Description)
}
