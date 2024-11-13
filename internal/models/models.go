package models

type Qualifiers struct {
	aptDate             int64
	patientAge          uint8
	numberOfBodySystems uint8
	aptDuration         uint8
	practiceType        string
	isAptVirtual        string
	attendeesPresent    string
	primaryIssue        string
	secondaryIssue      string
	isRecentDischarge   string
	formsCompleted      string
	procedureType       string
}
