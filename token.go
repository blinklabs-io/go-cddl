package cddl

import "github.com/bzick/tokenizer"

type TokenType tokenizer.TokenKey

const (
	TokenUndefined TokenType = iota

	TokenControlOperatorSize

/*
  /// Identifier with optional `SocketPlug`
  IDENT((&'a str, Option<SocketPlug>)),
  /// Value
  VALUE(Value<'a>),
  /// CBOR tag '#'
  TAG((Option<u8>, Option<usize>)),

  // Operators
  /// Assignment operator '='
  ASSIGN,
  /// Optional occurrence indicator '?'
  OPTIONAL,
  /// Zero or more occurrence indicator '*'
  ASTERISK,
  /// One or more occurrence indicator '+'
  ONEORMORE,
  /// Unwrap operator '~'
  UNWRAP,

  // Delimiters
  /// Comma ','
  COMMA,
  /// Colon ':'
  COLON,

  /// Comment text
  COMMENT(&'a str),

  /// Type choice indicator '/'
  TCHOICE,
  /// Group choice indicator '//'
  GCHOICE,
  /// Type choice alternative '/='
  TCHOICEALT,
  /// Group choice alternative '//='
  GCHOICEALT,
  /// Arrow map '=>'
  ARROWMAP,
  /// Cut '^'
  CUT,

  /// Range operator. Inclusive '..' if true, otherwise exclusive '...'s
  RANGEOP(bool),

  /// Range tuple with lower bound, upper bound, and bool indicating whether or
  /// not the range is inclusive
  RANGE((RangeValue<'a>, RangeValue<'a>, bool)),

  /// Left opening parend
  LPAREN,
  /// Right closing parend
  RPAREN,
  /// Left opening brace
  LBRACE,
  /// Right closing brace
  RBRACE,
  /// Left opening bracket
  LBRACKET,
  /// Right closing bracket
  RBRACKET,
  /// Left opening angle bracket
  LANGLEBRACKET,
  /// Right closing angle bracket
  RANGLEBRACKET,

  // Control operators
  /// .size control operator
  SIZE,
  /// .bits control operator
  BITS,
  /// .regexp control operator
  CREGEXP,
  /// .cbor control operator
  CBOR,
  /// .cborseq control operator
  CBORSEQ,
  /// .within control operator
  WITHIN,
  /// .and control operator
  AND,
  /// .lt control operator
  LT,
  /// .le control operator
  LE,
  /// .gt control operator
  GT,
  /// .ge control operator
  GE,
  /// .eq control operator
  EQ,
  /// .ne control operator
  NE,
  /// .default control operator
  DEFAULT,
  /// .pcre control operator
  /// Proposed control extension to support Perl-Compatible Regular Expressions
  /// (PCREs). See https://tools.ietf.org/html/rfc8610#section-3.8.3.2s
  PCRE,

  /// group to choice enumeration '&'
  GTOCHOICE,

  // Standard prelude
  // https://www.rfc-editor.org/rfc/rfc8610.html#appendix-D
  /// false
  FALSE,
  /// true
  TRUE,
  /// bool
  BOOL,
  /// nil
  NIL,
  /// null
  NULL,
  /// uint
  UINT,
  /// nint
  NINT,
  /// int
  INT,
  /// float16
  FLOAT16,
  /// float32
  FLOAT32,
  /// float64
  FLOAT64,
  /// float16-32
  FLOAT1632,
  /// float32-64
  FLOAT3264,
  ///float
  FLOAT,
  /// bstr
  BSTR,
  /// tstr
  TSTR,
  /// any
  ANY,
  /// bytes
  BYTES,
  /// text
  TEXT,
  /// tdate
  TDATE,
  /// time
  TIME,
  /// number
  NUMBER,
  /// biguint
  BIGUINT,
  /// bignint
  BIGNINT,
  /// bigint
  BIGINT,
  /// integer
  INTEGER,
  /// unsigned
  UNSIGNED,
  /// decfrac
  DECFRAC,
  /// bigfloat
  BIGFLOAT,
  /// eb64url
  EB64URL,
  /// eb64legacy
  EB64LEGACY,
  /// eb16k
  EB16,
  /// encoded-cbor
  ENCODEDCBOR,
  /// uri
  URI,
  /// b64url
  B64URL,
  /// b64legacy
  B64LEGACY,
  /// regexp
  REGEXP,
  /// mime-message
  MIMEMESSAGE,
  /// cbor-any
  CBORANY,
  /// undefined
  UNDEFINED,
*/
)

var tokenTypeNames = map[TokenType]string{
	// Types from tokenizer library
	TokenType(tokenizer.TokenUnknown): "UNKNOWN",
	TokenType(tokenizer.TokenString):  "STRING",
	TokenType(tokenizer.TokenFloat):   "FLOAT",
	TokenType(tokenizer.TokenInteger): "INTEGER",
	TokenType(tokenizer.TokenKeyword): "KEYWORD",
	// Custom types
	TokenControlOperatorSize: "SIZE",
}

func (t TokenType) String() string {
	name, ok := tokenTypeNames[t]
	if !ok {
		return "UNKNOWN"
	}
	return name
}
