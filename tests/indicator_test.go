package test

import (
	"testing"

	"github.com/DaanWillems/stix2-go/stix"
)

func TestIndicator(t *testing.T) {
	i := stix.NewIndicator(
		"aaa",
		stix.CommonSDOOptions(
			stix.WithConfidence(80),
			stix.WithLang("English"),
			stix.WithLabels([]string{"aa, bb"}),
		),
		[]stix.IndicatorOption{
			stix.WithIndicatorName("aa"),
			// stix.WithIndicatorDescription("bb"),
			stix.WithPatternType("stix"),
			stix.WithPattern("wdwdwd"),
		},
	)

	if err := i.Validate(); err != nil {
		t.Errorf("Struct does not validate.: %#v", err.Error())
	}
}
