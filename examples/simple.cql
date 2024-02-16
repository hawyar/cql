library SomeOrg.SimpleLib.Anotha1 version '1.0'

include CMS153_Common version '1.3' called Common
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


// code
// code "Hepatitis Type B (SNOMED)": '66071002' from "SNOMED-CT" display 'Viral hepatitis type B (disorder)'
// code "Hepatitis Type B (ICD-10)": 'B18.1' from "ICD-10-CM" display 'Chronic viral hepatitis B without delta-agent'


// define statement
define "MalePatient20To40":
  AgeInYearsAt(start of MeasurementPeriod) >= 20
    and AgeInYearsAt(start of MeasurementPeriod) < 40
    and Patient.gender in "Male Administrative Sex"

define function ToCalendarUnit(unit System.String):
    case unit
        when 'ms' then 'millisecond'
        when 's' then 'second'
        when 'min' then 'minute'
        when 'h' then 'hour'
        when 'd' then 'day'
        when 'wk' then 'week'
        when 'mo' then 'month'
        when 'a' then 'year'
        else unit
    end