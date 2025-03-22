package main

import (
	"testing"
	"time"
)

func TestIndicator(t *testing.T) {
	indicator := NewIndicator(
		CommonSDOOptions(
			WithSDOConfidence(80),
			WithSDOLang("English"),
			WithSDOLabels([]string{"aa, bb"}),
		),
		[]IndicatorOption{
			WithIndicatorName("aa"),
			WithIndicatorDescription("bb"),
			WithIndicatorPatternType("stix"),
			WithIndicatorPattern("wdwdwd"),
			WithIndicatorValidFrom(time.Date(2025, 3, 7, 0, 0, 0, 0, time.UTC)),
			WithIndicatorTypes([]string{"ipv4"}),
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
	threatActor := NewThreatActor(
		CommonSDOOptions(
			WithSDOSpecVersion("2.1"),
			WithSDOConfidence(90),
			WithSDOLabels([]string{"apt", "state-sponsored"}),
		),
		[]ThreatActorOption{
			WithThreatActorName("APT-42"),
			WithThreatActorDescription("Advanced threat actor targeting financial institutions"),
			WithThreatActorTypes([]string{"Big hacker man"}),
			WithThreatActorAliases([]string{"FinancialWizard", "MoneyHunters"}),
			WithThreatActorFirstSeen(time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC)),
			WithThreatActorRoles([]string{"agent", "infiltrator"}),
			WithThreatActorSophistication("advanced"),
			WithThreatActorResourceLevel("government"),
			WithThreatActorPrimaryMotivation("financial-gain"),
			WithThreatActorSecondaryMotivation([]string{"intelligence-gathering", "dominance"}),
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
	grouping := NewGrouping(
		CommonSDOOptions(
			WithSDOSpecVersion("2.1"),
			WithSDOConfidence(90),
			WithSDOLabels([]string{"apt", "state-sponsored"}),
		),
		[]GroupingOption{
			WithGroupingName("Group"),
			WithGroupingDescription("Group Desc"),
			WithGroupingContext([]string{
				"a",
				"b",
			}),
			WithGroupingObjectRefs([]string{
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

	grouping2 := NewGrouping(
		CommonSDOOptions(
			WithSDOSpecVersion("2.1"),
			WithSDOConfidence(90),
			WithSDOLabels([]string{"apt", "state-sponsored"}),
		),
		[]GroupingOption{
			WithGroupingName("Group2"),
			WithGroupingDescription("Group2 Desc"),
			WithGroupingContext([]string{
				"a",
				"b",
			}),
			WithGroupingObjectRefs([]string{
				"a",
			}),
		},
	)

	grouping2.ID = grouping2.GenerateID()

	if grouping.ID == grouping2.ID {
		t.Error("ID's should be unique")
	}

}
