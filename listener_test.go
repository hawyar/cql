package cql

// func TestParse(t *testing.T) {
// 	t.Run("LibraryDefinition", func(t *testing.T) {

// 		t.Run("Identifier", func(t *testing.T) {
// 			libDefinitions := map[string]Library{
// 				"library CQM_Mod": {
// 					LibraryDefinition: LibraryDefinition{
// 						QualifiedIdentifier: QualifiedIdentifier{
// 							Identifier: "CQM_Mod",
// 						},
// 					},
// 				},
// 				"library \"CQM_Mod\"": {
// 					LibraryDefinition: LibraryDefinition{
// 						QualifiedIdentifier: QualifiedIdentifier{
// 							Identifier: "\"CQM_Mod\"",
// 						},
// 					},
// 				},
// 				"library MyLib.`SubLib`.\"Quoted\".Common version '1.0.0' using FHIR version '4.0.1'": {
// 					LibraryDefinition: LibraryDefinition{
// 						QualifiedIdentifier: QualifiedIdentifier{
// 							Identifier: "Common",
// 							Qualifiers: []string{"MyLib", "`SubLib`", "\"Quoted\""},
// 						},
// 						Version: "1.0.0",
// 					},
// 					UsingDefinitions: []UsingDefinition{
// 						{
// 							Identifier: "FHIR",
// 							Version:    "4.0.1",
// 						},
// 					},
// 				},
// 			}

// 			for input, expected := range libDefinitions {
// 				stream := antlr.NewInputStream(input)

// 				ast, err := Parse(stream)

// 				if err != nil {
// 					t.Error(err)
// 				}

// 				assertEqual(t, expected, ast.Library)
// 			}
// 		})

// 		t.Run("IdentifierAndVersion", func(t *testing.T) {

// 			withVersion := map[string]Library{
// 				"library CQM_Mod version '1.0.0'": {
// 					LibraryDefinition: LibraryDefinition{
// 						QualifiedIdentifier: QualifiedIdentifier{
// 							Identifier: "CQM_Mod",
// 						},
// 						Version: "1.0.0",
// 					},
// 				},
// 			}

// 			for input, expected := range withVersion {
// 				stream := antlr.NewInputStream(input)

// 				ast, err := Parse(stream)

// 				if err != nil {
// 					t.Error(err)
// 				}

// 				assertEqual(t, expected, ast.Library)
// 			}
// 		})
// 	})

// 	t.Run("IncludeDefinition", func(t *testing.T) {
// 		includeDef := map[string]Library{
// 			"library Common version '1.0.0' using FHIR version '4.0.1'": {
// 				LibraryDefinition: LibraryDefinition{
// 					QualifiedIdentifier: QualifiedIdentifier{
// 						Identifier: "Common",
// 					},
// 					Version: "1.0.0",
// 				},
// 				UsingDefinitions: []UsingDefinition{
// 					{
// 						Identifier: "FHIR",
// 						Version:    "4.0.1",
// 					},
// 				},
// 			},
// 		}

// 		for input, expected := range includeDef {
// 			stream := antlr.NewInputStream(input)

// 			ast, err := Parse(stream)

// 			if err != nil {
// 				t.Error(err)
// 			}

// 			assertEqual(t, expected, ast.Library)
// 		}
// 	})
// }

// func assertEqual(t *testing.T, expected, actual Library) {
// 	if !reflect.DeepEqual(expected, actual) {
// 		t.Errorf("expected: %v \n got: %v", expected, actual)
// 	}
// }
