library ChlamydiaScreening_CQM version '2'

using QUICK version '1.0.0'

include CMS153_Common.Pkg version '1.3' called Common

public valueset "Female Administrative Sex": '2.16.840.1.113883.3.560.100.2'
valueset "Other Female Reproductive Conditions": '2.16.840.1.113883.3.464.1003.111.12.1006' version '1.0.0'
valueset "Genital Herpes": '2.16.840.1.113883.3.464.1003.110.12.1049'
valueset "Genococcal Infections and Venereal Diseases": '2.16.840.1.113883.3.464.1003.112.12.1001'
valueset "Inflammatory Diseases of Female Reproductive Organs": '2.16.840.1.113883.3.464.1003.112.12.1004'
valueset "Chlamydia": '2.16.840.1.113883.3.464.1003.112.12.1003'
valueset "HIV": '2.16.840.1.113883.3.464.1003.120.12.1003'
valueset "Syphilis": '2.16.840.1.113883.3.464.1003.112.12.1002'
valueset "Complications of Pregnancy, Childbirth and the Puerperium": '2.16.840.1.113883.3.464.1003.111.12.1012'
valueset "Pregnancy Test": '2.16.840.1.113883.3.464.1003.111.12.1011'
valueset "Pap Test": '2.16.840.1.113883.3.464.1003.108.12.1017'
valueset "Lab Tests During Pregnancy": '2.16.840.1.113883.3.464.1003.111.12.1007'
private valueset "Lab Tests for Sexually Transmitted Infections": '2.16.840.1.113883.3.464.1003.110.12.1051'
valueset "Chlamydia Screening": '2.16.840.1.113883.3.464.1003.110.12.1052' codesystems {
	code1, `DELIMITEDIDENTIFIER`, "doubleqouted"
}

parameter MeasurementPeriod default Interval[DateTime(2013, 1, 1, 0, 0, 0, 0), DateTime(2014, 1, 1, 0, 0, 0, 0))

context Patient
