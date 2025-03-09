package test

import (
	"testing"
	"time"

	"github.com/DaanWillems/stix2-go/stix"
)

func TestIndicator(t *testing.T) {
	indicator := stix.NewIndicator(
		"aaa",
		stix.CommonSDOOptions(
			stix.WithConfidence(80),
			stix.WithLang("English"),
			stix.WithLabels([]string{"aa, bb"}),
		),
		[]stix.IndicatorOption{
			stix.WithIndicatorName("aa"),
			stix.WithIndicatorDescription("bb"),
			stix.WithPatternType("stix"),
			stix.WithPattern("wdwdwd"),
			stix.WithValidFrom(time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC)),
			stix.WithIndicatorTypes([]string{"ipv4"}),
		},
	)

	if err := indicator.Validate(); err != nil {
		t.Errorf("Struct does not validate.: %#v", err.Error())
	}
}

func TestThreatActor(t *testing.T) {
	threatActor := stix.NewThreatActor(
		"threat-actor--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
		stix.CommonSDOOptions(
			stix.WithSpecVersion("2.1"),
			stix.WithConfidence(90),
			stix.WithLabels([]string{"apt", "state-sponsored"}),
		),
		[]stix.ThreatActorOption{
			stix.WithThreatActorName("APT-42"),
			stix.WithThreatActorDescription("Advanced threat actor targeting financial institutions"),
			stix.WithThreatActorTypes([]string{"Big hacker man"}),
			stix.WithAliases([]string{"FinancialWizard", "MoneyHunters"}),
			stix.WithFirstSeen(time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC)),
			stix.WithRoles([]string{"agent", "infiltrator"}),
			stix.WithSophistication("advanced"),
			stix.WithResourceLevel("government"),
			stix.WithPrimaryMotivation("financial-gain"),
			stix.WithSecondaryMotivation([]string{"intelligence-gathering", "dominance"}),
		},
	)

	if err := threatActor.Validate(); err != nil {
		t.Errorf("Struct does not validate.: %#v", err.Error())
	}
}

func TestGrouping(t *testing.T) {
	grouping := stix.NewGrouping(
		"threat-actor--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
		stix.CommonSDOOptions(
			stix.WithSpecVersion("2.1"),
			stix.WithConfidence(90),
			stix.WithLabels([]string{"apt", "state-sponsored"}),
		),
		[]stix.GroupingOption{
			stix.WithContext([]string{
				"a",
				"b",
			}),
			stix.WithObjectRefs([]string{
				"a",
			}),
		},
	)

	if err := grouping.Validate(); err != nil {
		t.Errorf("Struct does not validate.: %#v", err.Error())
	}
}
