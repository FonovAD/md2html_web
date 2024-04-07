package MarkdownToHTML

import (
	"bytes"
	"md2html_web/pkg/md2html/processing"
	"strings"
)

func Convert(MdText string) (string, error) {
	MdTextSplited := strings.Split(MdText, "\n")
	var HTML bytes.Buffer
	HTML.Grow(len(MdText) + len(HTMLPrefix) + len(HTMLPostfix))
	HTML.WriteString(HTMLPrefix)
	for _, j := range MdTextSplited {
		lex := processing.Lexer{
			Code:      string(j),
			Pos:       0,
			TokenList: []processing.Token{},
		}
		if err := lex.LexAnalusis(); err != nil {
			return "", err
		}
		parser := processing.Parser{Tokens: lex.TokenList, Pos: 0}
		root := parser.NewParseText()
		HTMLsize := (len(j) * HTMLsizeMultiplier) / HTMLsizeDevisor
		HTML.WriteString(processing.Run(root, HTMLsize))
	}
	HTML.WriteString(HTMLPostfix)
	return HTML.String(), nil
}
