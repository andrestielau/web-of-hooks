//go:build tools

package _package

import (
	_ "github.com/a-h/templ/cmd/templ"
	_ "github.com/cosmtrek/air"
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/jschaf/pggen/cmd/pggen"
	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)
