library "Hypertension Blood Pressure Control"

using QUICK

codesystem "LOINC": 'http://loinc.org'

valueset "Systolic Blood Pressure": '8480-6'
valueset "Diastolic Blood Pressure": '8462-4'
valueset "Hypertension Diagnosis": '38341003'

parameter "Measurement Period" Interval<DateTime>

context Patient

define "Hypertensive Patients":
  exists ("Hypertension Diagnosis" HypertensionDx
    where HypertensionDx.prevalenceRange overlaps "Measurement Period"
  )

define "Controlled Blood Pressure":
  Last(["Systolic Blood Pressure" SystolicBP, "Diastolic Blood Pressure" DiastolicBP])
    where SystolicBP.effective measurementPeriod.ends during "Measurement Period"
      and DiastolicBP.effective measurementPeriod.ends during "Measurement Period"
      and SystolicBP.value < 130
      and DiastolicBP.value < 80

define "Numerator":
  "Hypertensive Patients"
    and "Controlled Blood Pressure"

define "Denominator":
  "Hypertensive Patients"

define "Hypertension Blood Pressure Control Rate":
  Count("Numerator") / Count("Denominator")