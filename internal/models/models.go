package models

type Qualifiers struct {
	AppointmentDate         int64  // Timestamp of the appointment
	PatientAge              uint8  // Age of the patient
	NumberOfBodySystems     uint8  // Number of body systems evaluated
	AppointmentDuration_min uint8  // Duration of the appointment in minutes
	PracticeType            string // Type of medical practice (e.g., FHO, Clinic, ER)
	AppointmentType         string // Type of appointment (e.g., follow-up, consultation)
	AttendeesPresent        string // List of attendees present at the appointment
	PrimaryIssue            string // The primary medical issue being addressed
	SecondaryIssue          string // Any secondary medical issues being addressed
	IsRecentDischarge       bool   // Whether the patient was recently discharged from the hospital
	FormsCompleted          string // Whether forms were completed before or during the appointment
	ProcedureType           string // Type of procedure performed (if applicable)
}
