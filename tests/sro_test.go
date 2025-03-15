package test

import (
	"testing"
	"time"

	"github.com/DaanWillems/stix2-go/stix"
)

func TestRelationship(t *testing.T) {
	i := stix.NewRelationship(
		stix.CommonSROOptions(
			stix.WithSROConfidence(80),
			stix.WithSROLang("English"),
			stix.WithSROLabels([]string{"aa, bb"}),
		),
		[]stix.RelationshipOption{
			stix.WithRelationshipType("based-on"),
			stix.WithRelationshipDescription("aa"),
			stix.WithSourceRef("stix"),
			stix.WithTargetRef("stix"),
			stix.WithStartTime(time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC)),
			stix.WithStopTime(time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC)),
		},
	)

	i.ID = i.GenerateID()

	if err := i.Validate(); err != nil {
		t.Errorf("Struct does not validate.: %#v", err.Error())
	}
}
