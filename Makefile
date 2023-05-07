gen:
	@echo "generating parser..."
	antlr -Dlanguage=Go -o parser ./grammar/cql.g4  