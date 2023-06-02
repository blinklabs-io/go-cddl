package cddl

import (
	"github.com/alecthomas/participle/v2/lexer"
)

type AstDocument struct {
	Pos         lexer.Position
	Definitions []*AstDefinition `parser:"@@"`
}

type AstDefinition struct {
	Pos         lexer.Position
	Identifier  string   `parser:"@Ident '='"`
	Expressions []string `parser:"@Expression+"`
}
