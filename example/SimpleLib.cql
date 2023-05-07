library SimpleLib

using FHIR version '4.0.1'

context Patient

define isAdult:
    age >= 18

define hasDiagnosis:
    condition.exists()

define isHighRisk:
    isAdult and hasDiagnosis
