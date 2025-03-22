package stix

//go:generate ./bin/generator $GOFILE sdo
type ExternalReference struct {
	SDO
	SourceName string `json:"source_name"`
	ExternalID string `json:"external_id"`
}
