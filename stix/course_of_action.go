package stix

type CourseOfAction struct {
	SDO
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	ActionType      []string          `json:"os_execution_envs"`
	OsExecutionEnvs []string          `json:"os_execution_envs "`
	ActionBin       string            `json:"action_bin "`
	ActionReference ExternalReference `json:"action_reference "`
}
