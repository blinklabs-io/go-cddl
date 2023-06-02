package cddl

import "github.com/alecthomas/participle/v2/lexer"

var (
	parserLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Whitespace", `[ \t\n\r]+`},
		{"Comment", `(?://)[^\n]*\n?`},
		{"Ident", `[@_$a-z][-@_$a-z0-9]*`},
		{"Equal", `=`},
		{"Expression", `[\w\W]+`},
		/*
			{"Ident", `[a-zA-Z]\w*`, nil},
			{"Number", `(?:\d*\.)?\d+`, nil},
			{"Punct", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
			{"Whitespace", `[ \t\n\r]+`, nil},
		*/
	})
)
