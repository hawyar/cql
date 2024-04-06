package cql

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/hawyar/cql/internal"
)

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

type CQLListener struct {
	*parser.BasecqlListener
	ast *AST
}

func NewCQLListener() *CQLListener {
	return &CQLListener{
		ast: &AST{},
	}
}

func (c *CQLListener) AST() *AST {
	return c.ast
}

func ConsumeIdentifier(ctx *parser.IdentifierContext) Identifier {
	if ctx == nil {
		return nil
	}

	if ctx.QUOTEDIDENTIFIER() != nil {
		return QoutedIdentifier{
			Name: ctx.QUOTEDIDENTIFIER().GetText(),
		}
	}

	if ctx.DELIMITEDIDENTIFIER() != nil {
		return DelimtedIdentifier{
			Name: ctx.DELIMITEDIDENTIFIER().GetText(),
		}
	}

	if ctx.IDENTIFIER() != nil {
		return PlainIdentifier{
			Name: ctx.IDENTIFIER().GetText(),
		}
	}

	return nil
}

func ConsumeQualifier(ctx *parser.QualifierContext) Identifier {
	if ctx.Identifier().IDENTIFIER() != nil {
		return PlainIdentifier{
			Name: ctx.Identifier().IDENTIFIER().GetText(),
		}
	}

	if ctx.Identifier().QUOTEDIDENTIFIER() != nil {
		return QoutedIdentifier{
			Name: ctx.Identifier().QUOTEDIDENTIFIER().GetText(),
		}
	}

	if ctx.Identifier().DELIMITEDIDENTIFIER() != nil {
		return DelimtedIdentifier{
			Name: ctx.Identifier().DELIMITEDIDENTIFIER().GetText(),
		}
	}

	return nil
}

func ConsumeQualifiedIdentifier(ctx *parser.QualifiedIdentifierContext) QualifiedIdentifier {
	ident := QualifiedIdentifier{
		Resolved: ctx.GetText(),
	}

	if ctx.Identifier() != nil {
		ident.Identifier = ConsumeIdentifier(ctx.Identifier().(*parser.IdentifierContext))
	}

	if ctx.AllQualifier() != nil {
		for _, q := range ctx.AllQualifier() {
			ident.Qualifiers = append(ident.Qualifiers, ConsumeQualifier(q.(*parser.QualifierContext)))
		}
	}

	return ident
}

func (c *CQLListener) EnterLibrary(ctx *parser.LibraryContext) {
	if ctx.LibraryDefinition() != nil {
		lib := LibraryDefinition{}

		if ctx.LibraryDefinition().QualifiedIdentifier() != nil {
			lib.QualifiedIdentifier = ConsumeQualifiedIdentifier(ctx.LibraryDefinition().QualifiedIdentifier().(*parser.QualifiedIdentifierContext))
		}

		if ctx.LibraryDefinition().VersionSpecifier() != nil {
			lib.Version = ctx.LibraryDefinition().VersionSpecifier().GetText()
		}

		c.ast.Library = Library{
			LibraryDefinition: lib,
		}
	}

	if ctx.AllUsingDefinition() != nil {
		for _, def := range ctx.AllUsingDefinition() {
			using := UsingDefinition{}

			// ConsumeIdentifier would do just fine here but we need a qualifier as per the grammar
			if def.ModelIdentifier() != nil {
				if def.ModelIdentifier().Identifier().DELIMITEDIDENTIFIER() != nil {
					using.Identifier = DelimtedIdentifier{
						Name: def.ModelIdentifier().Identifier().DELIMITEDIDENTIFIER().GetText(),
					}
				}
			}

			if def.VersionSpecifier() != nil {
				using.Version = strings.Trim(def.VersionSpecifier().GetText(), "'")
			}

			c.ast.Library.UsingDefinitions = append(c.ast.Library.UsingDefinitions, using)
		}
	}

	if ctx.AllIncludeDefinition() != nil {
		for _, include := range ctx.AllIncludeDefinition() {
			def := IncludeDefinition{}
			if include.QualifiedIdentifier() != nil {
				def.QualifiedIdentifier = ConsumeQualifiedIdentifier(include.QualifiedIdentifier().(*parser.QualifiedIdentifierContext))
			}

			if include.VersionSpecifier() != nil {
				def.Version = include.VersionSpecifier().GetText()
			}

			if include.LocalIdentifier() != nil {
				if include.LocalIdentifier().Identifier().IDENTIFIER() != nil {
					def.Alias = DelimtedIdentifier{
						Name: include.LocalIdentifier().Identifier().IDENTIFIER().GetText(),
					}
				} else if include.LocalIdentifier().Identifier().DELIMITEDIDENTIFIER() != nil {
					def.Alias = DelimtedIdentifier{
						Name: include.LocalIdentifier().Identifier().DELIMITEDIDENTIFIER().GetText(),
					}
				} else if include.LocalIdentifier().Identifier().QUOTEDIDENTIFIER() != nil {
					def.Alias = QoutedIdentifier{
						Name: include.LocalIdentifier().Identifier().QUOTEDIDENTIFIER().GetText(),
					}
				}
			}
			c.ast.Library.IncludeDefinitions = append(c.ast.Library.IncludeDefinitions, def)
		}

		if ctx.AllCodesystemDefinition() != nil {
			for _, codesystem := range ctx.AllCodesystemDefinition() {
				def := CodesystemDefinition{}

				if codesystem.AccessModifier() != nil {
					def.AccessModifier = AccessModifier(strings.Trim(codesystem.AccessModifier().GetText(), "\""))
				}

				if codesystem.Identifier() != nil {
					def.Identifier = ConsumeIdentifier(codesystem.Identifier().(*parser.IdentifierContext))
				}

				if codesystem.VersionSpecifier() != nil {
					def.Version = strings.Trim(codesystem.VersionSpecifier().GetText(), "\"")
				}

				if codesystem.CodesystemId() != nil {
					def.CodesystemID = strings.Trim(codesystem.CodesystemId().GetText(), "'")
				}

				c.ast.Library.CodesystemDefinitions = append(c.ast.Library.CodesystemDefinitions, def)
			}
		}
	}

	if ctx.AllValuesetDefinition() != nil {
		valuesets := make([]ValuesetDefinition, 0)
		for _, v := range ctx.AllValuesetDefinition() {
			def := ValuesetDefinition{}

			if v.AccessModifier() != nil {
				def.AccessModifier = AccessModifier(strings.Trim(v.AccessModifier().GetText(), "\""))
			}

			if v.Identifier() != nil {
				def.Identifier = ConsumeIdentifier(v.Identifier().(*parser.IdentifierContext))
			}

			if v.VersionSpecifier() != nil {
				def.Version = strings.Trim(v.VersionSpecifier().GetText(), "\"")
			}

			if v.ValuesetId() != nil {
				def.ValuesetID = strings.Trim(v.ValuesetId().GetText(), "'")
			}

			if v.Codesystems() != nil {
				codesystem := Codesystems{
					QualifiedIdentifier: make([]QualifiedIdentifier, 0),
				}

				for _, id := range v.Codesystems().AllCodesystemIdentifier() {
					out := ConsumeIdentifier(id.Identifier().(*parser.IdentifierContext))

					if out != nil {
						codesystem.QualifiedIdentifier = append(codesystem.QualifiedIdentifier, QualifiedIdentifier{
							Identifier: out,
						})
					}
				}

				def.Codesystems = codesystem
			}

			valuesets = append(valuesets, def)
		}
		c.ast.Library.ValuesetDefinitions = valuesets
	}

	if ctx.AllCodeDefinition() != nil {
		c.ast.Library.CodeDefinitions = make([]CodeDefinition, 0)

		for _, code := range ctx.AllCodeDefinition() {
			def := CodeDefinition{}

			if code.AccessModifier() != nil {
				def.AccessModifier = AccessModifier(strings.Trim(code.AccessModifier().GetText(), "\""))
			}

			if code.Identifier() != nil {
				def.Identifier = ConsumeIdentifier(code.Identifier().(*parser.IdentifierContext))
			}

			if code.CodesystemIdentifier() != nil {
				def.QualifiedIdentifier = QualifiedIdentifier{
					Identifier: ConsumeIdentifier(code.CodesystemIdentifier().Identifier().(*parser.IdentifierContext)),
				}
			}

			if code.CodeId() != nil {
			}

			if code.DisplayClause() != nil {
				def.DisplayClause = DisplayClause(strings.Trim(code.DisplayClause().GetText(), "'"))
			}
			c.ast.Library.CodeDefinitions = append(c.ast.Library.CodeDefinitions, def)
		}
	}

	if ctx.AllParameterDefinition() != nil {
	}

	if ctx.AllConceptDefinition() != nil {
	}

	if ctx.AllStatement() != nil {
		for _, stmt := range ctx.AllStatement() {
			if stmt != nil {
				// if no context is specified then default to Unfiltered
				if stmt.ContextDefinition() != nil {
					// format: <qualifier>.<identifier> or <identifier>
					if stmt.ContextDefinition().Identifier() != nil {
						qi := QualifiedIdentifier{
							Resolved: stmt.ContextDefinition().GetText(),
						}

						if stmt.ContextDefinition().ModelIdentifier() != nil {
							// this is the qualifier
							if stmt.ContextDefinition().ModelIdentifier().Identifier() != nil {
								qi.Qualifiers = append(qi.Qualifiers, ConsumeIdentifier(stmt.ContextDefinition().ModelIdentifier().Identifier().(*parser.IdentifierContext)))
							}
						}

						if stmt.ContextDefinition().Identifier() != nil {
							qi.Identifier = ConsumeIdentifier(stmt.ContextDefinition().Identifier().(*parser.IdentifierContext))
						}

						c.ast.Library.ContextDefinition = ContextDefinition{
							QualifiedIdentifier: QualifiedIdentifier{
								Identifier: ConsumeIdentifier(stmt.ContextDefinition().Identifier().(*parser.IdentifierContext)),
							},
						}

						c.ast.Library.ContextDefinition.QualifiedIdentifier = qi

					} else {

						c.ast.Library.ContextDefinition = ContextDefinition{
							QualifiedIdentifier: QualifiedIdentifier{
								Identifier: PlainIdentifier{
									Name: "Unfiltered",
								},
							},
						}
					}

					if stmt.ExpressionDefinition() != nil {
					}

					if stmt.FunctionDefinition() != nil {
					}
				}
			}
		}
	}
}

type Err struct {
	Line, Column int
	Msg          string
}

func (c *Err) Error() string {
	return fmt.Sprintf("%s %d:%d", c.Msg, c.Line, c.Column)
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	synatxErrs []Err
	ambiguity  []Err
}

func (e *ErrorListener) SyntaxErrors() []Err {
	return e.synatxErrs
}

func (e *ErrorListener) SyntaxError(r antlr.Recognizer, d interface{}, line, column int, msg string, ex antlr.RecognitionException) {
	e.synatxErrs = append(e.synatxErrs, Err{line, column, msg})
}

func (e *ErrorListener) ReportAmbiguity(r antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (e *ErrorListener) ReportAttemptingFullContext(r antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
}

func (e *ErrorListener) ReportContextSensitivity(r antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
}

type Loc struct {
	Start int
	End   int
	Line  int
}

func (l Loc) Len() int {
	return l.End - l.Start
}

func (l Loc) String() string {
	return fmt.Sprintf("%d:%d %d", l.Start, l.End, l.Line)
}
