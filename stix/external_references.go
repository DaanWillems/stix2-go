package stix

type ExternalReference struct {
	SourceName string `json:"source_name"`
	ExternalID string `json:"external_id"`
}
