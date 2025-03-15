package test

import (
	"testing"
	"time"

	"github.com/DaanWillems/stix2-go/stix"
)

func TestIndicator(t *testing.T) {
	indicator := stix.NewIndicator(
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
	indicator.ID = indicator.GenerateID()

	if indicator.ID != "indicator--b1517d20-54bb-5706-b192-3823c4a68fa9" {
		t.Errorf("ID does not match expectation. Got: %v", indicator.ID)
	}

	if err := indicator.Validate(); err != nil {
		t.Errorf("Struct does not validate.: %#v", err.Error())
	}
}

func TestThreatActor(t *testing.T) {
	threatActor := stix.NewThreatActor(
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

	threatActor.ID = threatActor.GenerateID()

	if threatActor.ID != "threat-actor--19ef174d-85ca-5520-a7a4-d18c73f86c3a" {
		t.Errorf("ID does not match expectation. Got: %v", threatActor.ID)
	}

	if err := threatActor.Validate(); err != nil {
		t.Errorf("Struct does not validate.: %#v", err.Error())
	}
}

func TestGrouping(t *testing.T) {
	grouping := stix.NewGrouping(
		stix.CommonSDOOptions(
			stix.WithSpecVersion("2.1"),
			stix.WithConfidence(90),
			stix.WithLabels([]string{"apt", "state-sponsored"}),
		),
		[]stix.GroupingOption{
			stix.WithGroupingName("Group"),
			stix.WithGroupingDescription("Group Desc"),
			stix.WithContext([]string{
				"a",
				"b",
			}),
			stix.WithObjectRefs([]string{
				"a",
			}),
		},
	)

	grouping.ID = grouping.GenerateID()
	if grouping.ID != "grouping--18d8abf6-dc95-5359-8ed4-c65febe3250f" {
		t.Errorf("ID does not match expectation. Got: %v", grouping.ID)
	}
	if err := grouping.Validate(); err != nil {
		t.Errorf("Struct does not validate.: %#v", err.Error())
	}

	grouping2 := stix.NewGrouping(
		stix.CommonSDOOptions(
			stix.WithSpecVersion("2.1"),
			stix.WithConfidence(90),
			stix.WithLabels([]string{"apt", "state-sponsored"}),
		),
		[]stix.GroupingOption{
			stix.WithGroupingName("Group2"),
			stix.WithGroupingDescription("Group2 Desc"),
			stix.WithContext([]string{
				"a",
				"b",
			}),
			stix.WithObjectRefs([]string{
				"a",
			}),
		},
	)

	grouping2.ID = grouping2.GenerateID()

	if grouping.ID == grouping2.ID {
		t.Error("ID's should be unique")
	}

}
