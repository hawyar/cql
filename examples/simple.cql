library Ahealth.`HCL_Common` version '1.0'

using QUICK version '0.3.0'

include `CMS153_Common` version '1.3' called Common
include "CMS153_Common.Pkg"
include Basic version '1'

include FHIR

valueset "Female Administrative Sex": 'urn:oid:2.16.840.1.113883.3.560.100.2'
private valueset "Acute Pharyngitis": 'urn:oid:2.16.840.1.113883.3.464.1003.102.12.1011'

codesystems { "SNOMED-CT:2013-09", "ICD-9-CM:2014", "ICD-10-CM:2014" }

valueset "Diabetes": 'urn:oid:2.16.840.1.113883.3.464.1003.103.12.1001' version '20140501'

parameter MeasurementPeriod default Interval[@2013-01-01, @2014-01-01)
parameter Demographics Tuple { address String, city String, zip String }
parameter Points List<Integer>

context Bundle

