### Build

```bash
make build
```

That builds the `cql` and `fhirpath` executables in the `build` directory.

### Usage

Evaluate a CQL file

```bash
cql -input <path-to-cql-file>.cql -format xml
```

Evaluate a FHIRPath expression

```bash
fhirpath -e <fhirpath-expression>
```

### REPL

Both CQL and FHIRPath have a REPL mode.

```bash
cql repl
# or fhirpath repl
```

Spec:

- [CQL](http://cql.hl7.org/index.html)
- [FHIRPath Specification](http://hl7.org/fhirpath/)
