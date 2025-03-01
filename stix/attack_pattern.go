package stix

type AttackPattern struct {
	SDO
	Description     string   `json:"description"`
	Aliases         []string `json:"aliases"`
	KillChainPhases []string `json:"kill_chain_phases"`
}
