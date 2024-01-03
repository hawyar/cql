
library PopCalc version '1.0'

using QUICK version '0.3.0'
using AnotherLib

include CMS153_Common version '1.3' called Common

context Patient

define AdultPatient:
  age >= 18
