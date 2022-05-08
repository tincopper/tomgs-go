// Package generator provides an abstract interface to code generators.
package generator

import (
	"tomgs-go/learning-protoc/protoc-plugin/protoc-gen-dubbo-gateway/internal/descriptor"
)

// Generator is an abstraction of code generators.
type Generator interface {
	// Generate generates output files from input .proto files.
	Generate(targets []*descriptor.File) ([]*descriptor.ResponseFile, error)
}
