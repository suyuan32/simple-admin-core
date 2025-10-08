package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/chimerakang/simple-admin-core/tools/protoc-gen-go-zero-api/generator"
)

func main() {
	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		// Support Proto3 optional fields
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if err := generateFile(gen, f); err != nil {
				gen.Error(err)
				return err
			}
		}
		return nil
	})
}

// generateFile generates .api file for a single proto file
func generateFile(gen *protogen.Plugin, file *protogen.File) error {
	// Skip if no services defined
	if len(file.Services) == 0 {
		return nil
	}

	// Generate .api filename
	filename := getAPIFilename(file)
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	// Create generator instance
	apiGen := generator.NewGenerator(file)
	content, err := apiGen.Generate()
	if err != nil {
		return fmt.Errorf("failed to generate .api for %s: %w", file.Desc.Path(), err)
	}

	// Write generated content
	g.P(content)

	return nil
}

// getAPIFilename transforms proto path to .api path
// Example: rpc/desc/user.proto -> api/desc/core/user.api
func getAPIFilename(file *protogen.File) string {
	// Get base name without .proto extension
	name := file.GeneratedFilenamePrefix

	// Transform path: rpc/desc/xxx -> api/desc/core/xxx
	// For now, simple transformation - can be enhanced later
	apiPath := transformPath(name) + ".api"

	return apiPath
}

// transformPath transforms proto path to api path
// TODO: Make this configurable via plugin options
func transformPath(protoPath string) string {
	// Simple replacement for now
	// rpc/desc/user -> api/desc/core/user
	if len(protoPath) >= 8 && protoPath[:8] == "rpc/desc" {
		return "api/desc/core" + protoPath[8:]
	}
	return protoPath
}
