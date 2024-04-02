package cql

// Supported R4 FHIR resources, referenced in CQL expressions
var FHIRResources = map[string]string{
	"Unfiltered": "Unfiltered",
	// clinical (general)
	"AllergyIntolerance":  "AllergyIntolerance",
	"Condition":           "Condition",
	"Procedure":           "Procedure",
	"ClinicalImpression":  "ClinicalImpression",
	"FamilyMemberHistory": "FamilyMemberHistory",
	"RiskAssessment":      "RiskAssessment",
	"DetectedIssue":       "DetectedIssue",
	// care provision
	"CarePlan":           "CarePlan",
	"Goal":               "Goal",
	"ServiceRequest":     "ServiceRequest",
	"NutritionOrder":     "NutritionOrder",
	"NutritionIntake":    "NutritionIntake",
	"VisionPrescription": "VisionPrescription",
	// medication
	"Medication":                 "Medication",
	"MedicationKnowledge":        "MedicationKnowledge",
	"MedicationRequest":          "MedicationRequest",
	"MedicationAdministration":   "MedicationAdministration",
	"MedicationDispense":         "MedicationDispense",
	"MedicationStatement":        "MedicationStatement",
	"Immunization":               "Immunization",
	"ImmunizationEvaluation":     "ImmunizationEvaluation",
	"ImmunizationRecommendation": "ImmunizationRecommendation",
	// diagnostics
	"Observation":       "Observation",
	"DocumentReference": "DocumentReference",
	"DiagnosticReport":  "DiagnosticReport",
	"Specimen":          "Specimen",
	"BodyStructure":     "BodyStructure",
	"ImagingSelection":  "ImagingSelection",
	"ImagingStudy":      "ImagingStudy",
	"GenomicStudy":      "GenomicStudy",
	// individuals (identification)
	"Patient":       "Patient",
	"Practitioner":  "Practitioner",
	"RelatedPerson": "RelatedPerson",
	// groups
	"Organization":      "Organization",
	"HealthcareService": "HealthcareService",
	"Group":             "Group",
}

func SupportedResource(resource string) bool {
	_, ok := FHIRResources[resource]
	return ok
}
