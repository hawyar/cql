
// library definition
library Lib1.Ctx version '1.0.0'

// using (data models)
include CMS153_Common version '1.3' called Common
include FHIR

// valueset
valueset "Female Administrative Sex": 'urn:oid:2.16.840.1.113883.3.560.100.2'
valueset "Acute Pharyngitis": 'urn:oid:2.16.840.1.113883.3.464.1003.102.12.1011'

// parameters
parameter MeasurementPeriod default Interval[@2013-01-01, @2014-01-01)

// context
context Patient

// define statement
define "Patient16To23AndFemale":
  AgeInYearsAt(start of MeasurementPeriod) >= 16
    and AgeInYearsAt(start of MeasurementPeriod) < 24
    and Patient.gender in "Female Administrative Sex"
