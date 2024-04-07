package processing

type TokenType struct {
	name  string
	regex string
}

var TokenTypes = map[string]TokenType{
	// "PUNCTUATION":  TokenType{"PUNCTUATION", "\\D[.]{1}"},
	"SEMICOLON":    TokenType{"SEMICOLON", "[\\n|\\v]"},
	"HEADING":      TokenType{"HEADING", "[#]{1,6}"},
	"LINE":         TokenType{"LINE", "[=|-|*]{3,}"},
	"LIST":         TokenType{"LIST", "[*|+|-]{1}[ ]{1}"},
	"NUMBEREDLIST": TokenType{"NUMBEREDLIST", "\\d[.]"},
	"BOLT":         TokenType{"BOLT", "[*|_]{2}[\\w| ]{1,}[*|_]{2}"},
	"CODE":         TokenType{"CODE", "[`]"},
	"CODEBLOCK":    TokenType{"CODEBLOCK", "[`]{3}"},
	"WORD":         TokenType{"WORD", "\\w+[-]?[.|,|!|?]?"},
	"SPACE":        TokenType{"SPACE", "[ ]{1,}"},
	"ITALIC":       TokenType{"ITALIC", "[*|_]{1}[\\w| ]{1,}[*|_]{1}"},
	"SPECIALCHAR":  TokenType{"SPECIALCHAR", "[/|;|:|%|@|$|&|~]{1}"},
}

var SecondTokenTypes = map[string]TokenType{
	"BLOCKNUMBEREDLIST": TokenType{"BLOCKNUMBEREDLIST", "[]"},
}
