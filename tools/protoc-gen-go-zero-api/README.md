# protoc-gen-go-zero-api

A protoc plugin that generates Go-Zero `.api` files from Protocol Buffer definitions with `google.api.http` annotations.

## Overview

This plugin eliminates the need to maintain separate Proto and `.api` files by auto-generating Go-Zero API definitions from Proto services. It supports:

- ✅ HTTP method mapping (GET, POST, PUT, DELETE, PATCH)
- ✅ Path parameter transformation (`{id}` → `:id`)
- ✅ Go-Zero specific features (JWT, middleware, route groups) via custom Proto options
- ✅ Type conversion (Proto messages → Go-Zero types)
- 🚧 Service grouping by `@server` configuration (in progress)
- 🚧 Additional bindings support (in progress)

## Installation

```bash
# Build the plugin
cd tools/protoc-gen-go-zero-api
go build -o ../../bin/protoc-gen-go-zero-api

# Or use Makefile
make build-proto-plugin
```

## Usage

```bash
# Generate .api files from Proto
protoc --plugin=protoc-gen-go-zero-api=./bin/protoc-gen-go-zero-api \
       --go-zero-api_out=api/desc \
       --proto_path=. \
       rpc/desc/**/*.proto

# Or use Makefile
make gen-proto-api
```

## Project Structure

```
tools/protoc-gen-go-zero-api/
├── main.go              # Plugin entry point
├── generator/
│   ├── generator.go     # Main generator logic (stub)
│   ├── http_parser.go   # HTTP annotation parser (TODO: [PF-004])
│   ├── options_parser.go # Go-Zero options parser (TODO: [PF-006])
│   ├── type_converter.go # Type conversion (TODO: [PF-008])
│   ├── grouper.go       # Service grouping (TODO: [PF-010])
│   └── template.go      # .api template (TODO: [PF-012])
├── model/
│   ├── service.go       # Service model
│   ├── method.go        # Method model
│   └── message.go       # Message model
├── test/
│   ├── fixtures/        # Test Proto files
│   └── *_test.go        # Unit and integration tests (TODO)
├── go.mod
└── README.md
```

## Development Status

**Phase 1: Setup ✅ Completed** ([PF-001])
- ✅ Plugin project structure
- ✅ Go module initialization
- ✅ Main entry point
- ✅ Basic model definitions
- ✅ Stub generator
- ✅ Successful compilation

**Phase 2: Parsers** (TODO)
- 🚧 [PF-004] HTTP Annotation Parser
- 🚧 [PF-006] Options Parser
- 🚧 [PF-008] Type Converter

**Phase 3: Generation** (TODO)
- 🚧 [PF-010] Service Grouper
- 🚧 [PF-012] Template Generator
- 🚧 [PF-011] Integration

**Phase 4: Testing** (TODO)
- 🚧 [PF-013] Unit tests
- 🚧 [PF-014] Integration tests

## Current Limitations

- Currently generates stub `.api` files with basic info
- Full implementation in progress (see task-allocation.md)
- HTTP annotation parsing not yet implemented
- Go-Zero custom options not yet supported
- Type conversion not yet implemented

## Related Documentation

- **Specification**: `specs/003-proto-first-api-generation/spec.md`
- **Technical Plan**: `specs/003-proto-first-api-generation/plan.md`
- **Task Allocation**: `specs/003-proto-first-api-generation/task-allocation.md`

## Contributing

This is an internal tool for the Simple Admin Core project. See task allocation document for development assignments.

## License

Same as Simple Admin Core project.
