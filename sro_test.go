package main

import (
	"testing"
	"time"
)

func TestRelationship(t *testing.T) {
	i := NewRelationship(
		CommonSROOptions(
			WithSROConfidence(80),
			WithSROLang("English"),
			WithSROLabels([]string{"aa, bb"}),
		),
		[]RelationshipOption{
			WithRelationshipType("based-on"),
			WithRelationshipDescription("aa"),
			WithRelationshipSourceRef("stix"),
			WithRelationshipTargetRef("stix"),
			WithRelationshipStartTime(time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC)),
			WithRelationshipStopTime(time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC)),
		},
	)

	i.ID = i.GenerateID()

	if err := i.Validate(); err != nil {
		t.Errorf("Struct does not validate.: %#v", err.Error())
	}
}
