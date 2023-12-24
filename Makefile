BUILD_DIR = ./build
GRAMMAR_DIR = ./grammar
CQL_GRAMMAR = $(GRAMMAR_DIR)/cql.g4
CQL_PKG = parser # parser instead of cql to avoid name collision with the root package
FHIRPATH_GRAMMAR = $(GRAMMAR_DIR)/fhirpath.g4
FHIRPATH_PKG =fhirpath

OS=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)

.PHONY: all cql fhirpath build clean

all: cql fhirpath build clean

build: cql fhirpath
	@echo "Building CQL and FHIRPath executables..."
	GOOS=$(OS) GOARCH=$(ARCH) go build -o $(BUILD_DIR)/cql cmd/cql/main.go
	GOOS=$(OS) GOARCH=$(ARCH) go build -o $(BUILD_DIR)/fhirpath cmd/fhirpath/main.go

cql: has-antlr
	@echo "Generating CQL parser..."
	antlr -Dlanguage=Go -o $(CQL_PKG) -package $(CQL_PKG) -Xexact-output-dir $(CQL_GRAMMAR)

fhirpath: has-antlr
	@echo "Generating FHIRPath parser..."
	antlr -Dlanguage=Go -o $(FHIRPATH_PKG)/parser -package $(FHIRPATH_PKG) -Xexact-output-dir $(FHIRPATH_GRAMMAR)

has-antlr:
	@which antlr > /dev/null || (echo "Please install antlr4 and add it to your PATH" && exit 1)

clean:
	@echo "Cleaning up..."
	go clean
	rm -rf $(CQL_PKG) $(FHIRPATH_PKG)