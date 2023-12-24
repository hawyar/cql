all: cql fhirpath build

build:
	@echo "building..."

cql:
	@echo "generating CQL parser..."
	antlr -Dlanguage=Go -o parser ./grammar/fhirpath.g4 -package parser

fhirpath:
	@echo "generating FHIRPath parser..."
	antlr -Dlanguage=Go -o fhirpath ./grammar/fhirpath.g4  -package fhirpath

