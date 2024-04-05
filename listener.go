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

func ConsumeQualifiedIdentifier(ctx *parser.QualifiedIdentifierContext) QualifiedIdentifier {
	ident := QualifiedIdentifier{
		Resolved: ctx.GetText(),
		Loc: Loc{
			Start: ctx.GetStart().GetStart(),
			End:   ctx.GetStop().GetStop(),
			Line:  ctx.GetStop().GetLine(),
		},
	}

	if ctx.Identifier() != nil {
		if ctx.Identifier().IDENTIFIER() != nil {
			ident.Identifier = PlainIdentifier{
				Name: ctx.Identifier().IDENTIFIER().GetText(),
				Loc: Loc{
					Start: ctx.Identifier().IDENTIFIER().GetSymbol().GetStart(),
					End:   ctx.Identifier().IDENTIFIER().GetSymbol().GetStop(),
					Line:  ctx.Identifier().IDENTIFIER().GetSymbol().GetLine(),
				},
			}
		} else if ctx.Identifier().QUOTEDIDENTIFIER() != nil {
			ident.Identifier = QoutedIdentifier{
				Name: ctx.Identifier().QUOTEDIDENTIFIER().GetText(),
				Loc: Loc{
					Start: ctx.Identifier().QUOTEDIDENTIFIER().GetSymbol().GetStart(),
					End:   ctx.Identifier().QUOTEDIDENTIFIER().GetSymbol().GetStop(),
					Line:  ctx.Identifier().QUOTEDIDENTIFIER().GetSymbol().GetLine(),
				},
			}
		} else if ctx.Identifier().DELIMITEDIDENTIFIER() != nil {
			ident.Identifier = DelimtedIdentifier{
				Name: ctx.Identifier().DELIMITEDIDENTIFIER().GetText(),
				Loc: Loc{
					Start: ctx.Identifier().DELIMITEDIDENTIFIER().GetSymbol().GetStart(),
					End:   ctx.Identifier().DELIMITEDIDENTIFIER().GetSymbol().GetStop(),
					Line:  ctx.Identifier().DELIMITEDIDENTIFIER().GetSymbol().GetLine(),
				},
			}
		}
	}

	if ctx.AllQualifier() != nil {
		for _, q := range ctx.AllQualifier() {
			if q.Identifier().IDENTIFIER() != nil {
				ident.Qualifiers = append(ident.Qualifiers, PlainIdentifier{
					Name: q.Identifier().IDENTIFIER().GetText(),
					Loc: Loc{
						Start: q.Identifier().IDENTIFIER().GetSymbol().GetStart(),
						End:   q.Identifier().IDENTIFIER().GetSymbol().GetStop(),
					},
				})
			}

			if q.Identifier().QUOTEDIDENTIFIER() != nil {
				ident.Qualifiers = append(ident.Qualifiers, QoutedIdentifier{
					Name: q.Identifier().QUOTEDIDENTIFIER().GetText(),
					Loc: Loc{
						Start: q.Identifier().QUOTEDIDENTIFIER().GetSymbol().GetStart(),
						End:   q.Identifier().QUOTEDIDENTIFIER().GetSymbol().GetStop(),
					},
				})
			}

			if q.Identifier().DELIMITEDIDENTIFIER() != nil {
				ident.Qualifiers = append(ident.Qualifiers, DelimtedIdentifier{
					Name: q.Identifier().DELIMITEDIDENTIFIER().GetText(),
					Loc: Loc{
						Start: q.Identifier().DELIMITEDIDENTIFIER().GetSymbol().GetStart(),
						End:   q.Identifier().DELIMITEDIDENTIFIER().GetSymbol().GetStop(),
					},
				})
			}
		}
	}

	return ident
}

func (c *CQLListener) EnterLibrary(ctx *parser.LibraryContext) {
	if ctx.LibraryDefinition() != nil {
		lib := LibraryDefinition{}

		if ctx.LibraryDefinition().QualifiedIdentifier() != nil {
			qualifierCtx := ctx.LibraryDefinition().QualifiedIdentifier().(*parser.QualifiedIdentifierContext)
			lib.QualifiedIdentifier = ConsumeQualifiedIdentifier(qualifierCtx)
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
					delim := DelimtedIdentifier{
						Name: def.ModelIdentifier().Identifier().DELIMITEDIDENTIFIER().GetText(),
						Loc: Loc{
							Start: def.ModelIdentifier().Identifier().DELIMITEDIDENTIFIER().GetSymbol().GetStart(),
							End:   def.ModelIdentifier().Identifier().DELIMITEDIDENTIFIER().GetSymbol().GetStop(),
						},
					}
					using.Identifier = delim

					fmt.Println(using)
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
				qualifierCtx := include.QualifiedIdentifier().(*parser.QualifiedIdentifierContext)
				def.QualifiedIdentifier = ConsumeQualifiedIdentifier(qualifierCtx)
			}

			if include.VersionSpecifier() != nil {
				def.Version = include.VersionSpecifier().GetText()
			}

			if include.LocalIdentifier() != nil {
				if include.LocalIdentifier().Identifier().IDENTIFIER() != nil {
					ident := include.LocalIdentifier().Identifier().IDENTIFIER()
					def.Alias = DelimtedIdentifier{
						Name: ident.GetText(),
						Loc: Loc{
							Start: ident.GetSymbol().GetStart(),
							End:   ident.GetSymbol().GetStop(),
							Line:  ident.GetSymbol().GetLine(),
						},
					}
				} else if include.LocalIdentifier().Identifier().DELIMITEDIDENTIFIER() != nil {
					ident := include.LocalIdentifier().Identifier().DELIMITEDIDENTIFIER()
					def.Alias = DelimtedIdentifier{
						Name: ident.GetText(),
						Loc: Loc{
							Start: ident.GetSymbol().GetStart(),
							End:   ident.GetSymbol().GetStop(),
							Line:  ident.GetSymbol().GetLine(),
						},
					}
				} else if include.LocalIdentifier().Identifier().QUOTEDIDENTIFIER() != nil {
					ident := include.LocalIdentifier().Identifier().QUOTEDIDENTIFIER()
					def.Alias = QoutedIdentifier{
						Name: ident.GetText(),
						Loc: Loc{
							Start: ident.GetSymbol().GetStart(),
							End:   ident.GetSymbol().GetStop(),
							Line:  ident.GetSymbol().GetLine(),
						},
					}
				}
			}
			c.ast.Library.IncludeDefinitions = append(c.ast.Library.IncludeDefinitions, def)
		}

		if ctx.AllCodesystemDefinition() != nil {
		}
	}

	// for _, stmt := range ctx.AllStatement() {
	// 	if stmt != nil {
	// 		// if no context is specified in the library, and the model has not declared a default context
	// 		// 	// then set default context to Unfiltered.
	// 		if stmt.ContextDefinition() != nil {
	// 			fmt.Println("context", stmt.ContextDefinition().Identifier().IDENTIFIER().GetText())
	// 		}
	// 	}
	// }
}

// func (c *CQLListener) EnterValuesetDefinition(ctx *parser.ValuesetDefinitionContext) {
// 	c.stack.push(ctx)

// 	def := ValuesetDefinition{}

// 	if ctx.AccessModifier() != nil {
// 		fmt.Println("access modifier: ", ctx.AccessModifier().GetText())
// 		def.AccessModifier = strings.Trim(ctx.AccessModifier().GetText(), "\"")
// 	}

// 	if ctx.Identifier() != nil {
// 		fmt.Println("identifier: ", ctx.Identifier().GetText())
// 	}

// 	if ctx.VersionSpecifier() != nil {
// 		def.Version = strings.Trim(ctx.VersionSpecifier().GetText(), "\"")
// 	}

// 	if ctx.ValuesetId() != nil {
// 		// def.ID = strings.Trim(ctx.ValuesetId().GetText(), "\"")
// 		fmt.Println("valueset id: ", ctx.ValuesetId().GetText())
// 	}

// 	if ctx.Codesystems() != nil {
// 		codesystem := Codesystems{
// 			QualifiedIdentifier: make([]QualifiedIdentifier, 0),
// 		}

// 		// for _, id := range ctx.Codesystems().AllCodesystemIdentifier() {

// 		// 	codesystem.QualifiedIdentifier = append(codesystem.QualifiedIdentifier, qi)
// 		// }
// 		def.Codesystems = append(def.Codesystems, codesystem)
// 	}

// 	c.ast.Library.ValuesetDefinitions = append(c.ast.Library.ValuesetDefinitions, def)
// }

// func (c *CQLListener) ExitValuesetDefinition(ctx *parser.ValuesetDefinitionContext) {
// 	c.stack.pop()
// }

// // codesystems may appear without a valueset definition
// func (c *CQLListener) EnterCodesystems(ctx *parser.CodesystemsContext) {
// 	c.stack.push(ctx)

// 	codesystem := Codesystems{}

// 	// for _, id := range ctx.AllCodesystemIdentifier() {
// 	// 	qi := QualifiedIdentifier{
// 	// 		Identifier: strings.Trim(id.GetText(), "\""),
// 	// 	}
// 	// 	codesystem.QualifiedIdentifier = append(codesystem.QualifiedIdentifier, qi)
// 	// }

// 	c.ast.Library.CodesystemDefinitions = append(c.ast.Library.CodesystemDefinitions, codesystem)
// }

// func (c *CQLListener) ExitCodesystems(ctx *parser.CodesystemsContext) {
// 	c.stack.pop()
// }

// func (c *CQLListener) EnterParameterDefinition(ctx *parser.ParameterDefinitionContext) {
// 	c.stack.push(ctx)

// 	def := ParameterDefinition{}

// 	if ctx.AccessModifier() != nil {
// 		def.AccessModifier = AccessModifier(strings.Trim(ctx.AccessModifier().GetText(), "\""))
// 	}

// 	if ctx.Identifier() != nil {
// 		def.Name = strings.Trim(ctx.Identifier().GetText(), "\"")
// 	}

// 	if ctx.TypeSpecifier() != nil {
// 		// fmt.Println("type specifier: ", ctx.TypeSpecifier().GetText())
// 		def.TypeSpecifier = ctx.TypeSpecifier().GetText()

// 		if ctx.TypeSpecifier().NamedTypeSpecifier() != nil {
// 			fmt.Println("named type specifier: ", ctx.TypeSpecifier().NamedTypeSpecifier().GetText())
// 		}

// 		if ctx.TypeSpecifier().ListTypeSpecifier() != nil {
// 			fmt.Println("list type specifier: ", ctx.TypeSpecifier().ListTypeSpecifier().GetText())
// 		}

// 		if ctx.TypeSpecifier().IntervalTypeSpecifier() != nil {
// 			fmt.Println("interval type specifier: ", ctx.TypeSpecifier().IntervalTypeSpecifier().GetText())
// 		}

// 		if ctx.TypeSpecifier().TupleTypeSpecifier() != nil {
// 			// tuple := ctx.TypeSpecifier().TupleTypeSpecifier()
// 		}

// 		if ctx.TypeSpecifier().ChoiceTypeSpecifier() != nil {
// 			fmt.Println("choice type specifier: ", ctx.TypeSpecifier().ChoiceTypeSpecifier().GetText())
// 		}
// 	}

// 	// TODO: handle expression
// 	if ctx.Expression() != nil {
// 		// fmt.Println("expression: ", ctx.Expression().GetText())
// 	}

// 	c.ast.Library.ParamterDefinitions = append(c.ast.Library.ParamterDefinitions, def)
// }

// func (c *CQLListener) ExitParameterDefinition(ctx *parser.ParameterDefinitionContext) {
// 	c.stack.pop()
// }

// func (c *CQLListener) EnterCodeDefinition(ctx *parser.CodeDefinitionContext) {
// 	c.stack.push(ctx)

// 	def := CodeDefinition{}

// 	if ctx.AccessModifier() != nil {
// 		def.AccessModifier = AccessModifier(strings.Trim(ctx.AccessModifier().GetText(), "\""))
// 	}

// 	if ctx.Identifier() != nil {
// 		def.Identifier = strings.Trim(ctx.Identifier().GetText(), "\"")
// 	}

// 	if ctx.CodeId() != nil {
// 		def.CodeID = strings.Trim(ctx.CodeId().GetText(), "'")
// 	}

// 	if ctx.CodesystemIdentifier() != nil {
// 		// def.QualifiedIdentifier = QualifiedIdentifier{
// 		// 	Identifier: strings.Trim(ctx.CodesystemIdentifier().GetText(), "\""),
// 		// }
// 	}

// 	if ctx.DisplayClause() != nil {
// 		if strings.HasPrefix(ctx.DisplayClause().GetText(), "display") {
// 			def.DisplayClause = DisplayClause(strings.Trim(strings.Trim(ctx.DisplayClause().GetText(), "display"), "'"))
// 		}
// 	}

// 	c.ast.Library.CodeDefinitions = append(c.ast.Library.CodeDefinitions, def)
// }

// func (c *CQLListener) ExitCodeDefinition(ctx *parser.CodeDefinitionContext) {
// 	c.stack.pop()
// }

// func (c *CQLListener) EnterConceptDefinition(ctx *parser.ConceptDefinitionContext) {
// 	c.stack.push(ctx)

// 	concept := ConceptDefinition{}

// 	if ctx.AccessModifier() != nil {
// 		concept.AccessModifier = AccessModifier(strings.Trim(ctx.AccessModifier().GetText(), "\""))
// 	}

// 	if ctx.Identifier() != nil {
// 		concept.Identifier = strings.Trim(ctx.Identifier().GetText(), "\"")
// 	}

// 	if ctx.DisplayClause() != nil {
// 		if strings.HasPrefix(ctx.DisplayClause().GetText(), "display") {
// 			concept.DisplayClause = DisplayClause(strings.Trim(strings.Trim(ctx.DisplayClause().GetText(), "display"), "'"))
// 		}
// 	}

// 	// if ctx.AllCodeIdentifier() != nil {
// 	// 	for _, ident := range ctx.AllCodeIdentifier() {
// 	// 		qi := QualifiedIdentifier{
// 	// 			Identifier: strings.Trim(ident.GetText(), "\""),
// 	// 		}

// 	// 		concept.QualifiedIdentifier = append(concept.QualifiedIdentifier, qi)
// 	// 	}
// 	// }

// 	c.ast.Library.ConceptDefinitions = append(c.ast.Library.ConceptDefinitions, concept)
// }

// func (c *CQLListener) ExitConceptDefinition(ctx *parser.ConceptDefinitionContext) {
// 	c.stack.pop()
// }

// func (c *CQLListener) EnterExpressionDefinition(ctx *parser.ExpressionDefinitionContext) {
// 	// c.stack.push(ctx)
// 	fmt.Println("expression definition: ", ctx.GetText())
// }

type node struct {
	value antlr.ParserRuleContext
}

type stack struct {
	items []node
	size  int
}

func (s *stack) push(ctx antlr.ParserRuleContext) {
	s.items = append(s.items, node{value: ctx})
	s.size++
}

func (s *stack) pop() node {
	if s.size == 0 {
		return node{}
	}

	n := s.items[s.size-1]
	s.items = s.items[:s.size-1]
	s.size--

	return n
}

func (s *stack) peek() node {
	if s.size == 0 {
		return node{}
	}

	return s.items[s.size-1]
}

func (s *stack) isEmpty() bool {
	return s.size == 0
}

func (s *stack) clear() {
	s.items = []node{}
	s.size = 0
}

type Loc struct {
	Start int
	End   int
	Line  int
}

func (l Loc) Len() int {
	return l.Start - l.End
}

func (l Loc) String() string {
	return fmt.Sprintf("start: %d, end: %d, line: %d", l.Start, l.End, l.Line)
}
