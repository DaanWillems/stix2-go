package stix

//go:generate ./bin/generator $GOFILE
type AttackPattern struct {
	SDO
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Aliases         []string `json:"aliases"`
	KillChainPhases []string `json:"kill_chain_phases"`
}

func (AttackPattern *AttackPattern) GenerateID() string {
	return "AttackPattern--" + v5UUID(AttackPattern.Name+AttackPattern.Description)
}
