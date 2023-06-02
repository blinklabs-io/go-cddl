package cddl

import (
	"github.com/alecthomas/participle/v2"
)

type Parser struct {
	*participle.Parser[AstDocument]
}

func NewParser() *Parser {
	p := &Parser{
		Parser: participle.MustBuild[AstDocument](
			participle.Lexer(parserLexer),
			participle.Elide("Comment", "Whitespace"),
		),
	}
	return p
}
