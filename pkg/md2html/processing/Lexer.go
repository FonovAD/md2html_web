package processing

import (
	"regexp"
)

const (
	LexAnalusisDivisor = 2
	LexAnalusisError   = "I can't make out the words, try again"
)

type Lexer struct {
	Code      string
	Pos       int
	TokenList []Token
}

func (L *Lexer) LexAnalusis() error {
	count := 0
	for L.NextToken() {
		if count > len(L.Code)/LexAnalusisDivisor {
			return nil
		} else {
			count += 1
		}
	}
	return nil
}

func (L *Lexer) NextToken() bool {
	if L.Pos >= len(L.Code) {
		return false
	}
	var tokenTypesValues []string
	for key := range TokenTypes {
		tokenTypesValues = append(tokenTypesValues, key)
	}
	for i := 0; i < len(tokenTypesValues)-1; i++ {
		tokenType := TokenTypes[tokenTypesValues[i]]
		r, _ := regexp.Compile("^" + tokenType.regex)
		result := r.FindString(L.Code[L.Pos:])
		found := r.MatchString(L.Code[L.Pos:])
		if found {
			token := Token{tokenType, result}
			L.Pos = L.Pos + len(result)
			L.TokenList = append(L.TokenList, token)
			return true
		}
	}
	return true
}
