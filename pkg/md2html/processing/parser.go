package processing

import (
	"fmt"
)

var EmptyToken Token = Token{TokenType{"", ""}, ""}

var prefixHeadings = map[string]string{
	"#":      "<h1>",
	"##":     "<h2>",
	"###":    "<h3>",
	"####":   "<h4>",
	"#####":  "<h5>",
	"######": "<h6>",
}

var postfixHeadings = map[string]string{
	"#":      "</h1>",
	"##":     "</h2>",
	"###":    "</h3>",
	"####":   "</h4>",
	"#####":  "</h5>",
	"######": "</h6>",
}

type Parser struct {
	Tokens []Token
	Pos    int
}

func (P *Parser) Match(ExpectedTokenTypes []TokenType) Token {
	if P.Pos < len(P.Tokens) {
		currentToken := P.Tokens[P.Pos]
		for i := range ExpectedTokenTypes {
			if ExpectedTokenTypes[i].name == currentToken.Type.name {
				P.Pos += 1
				return currentToken
			}
		}
	}
	return Token{TokenType{"", ""}, ""}
}

func (P *Parser) Require(ExpectedTokenTypes []TokenType) Token {
	token := P.Match(ExpectedTokenTypes)
	emptyTokenType := TokenType{"", ""}
	if token.Type == emptyTokenType {
		errorStr := fmt.Sprintf("На позиции %d ожидается %s", P.Pos, ExpectedTokenTypes[0].name)
		panic(errorStr)
	}
	return token
}

func (P *Parser) NewParseText() StatmentsNode { //StatmentsNode - корень дерева(все узлы это строки)
	root := StatmentsNode{CodeString: []Node{}}
	for P.Pos < len(P.Tokens) {
		line := P.ParseLine()
		root.AddNode(line)
	}
	return root
}

func (P *Parser) ParseLine() Node {
	RecursiveTypes := []TokenType{
		TokenTypes["HEADING"],
		TokenTypes["NUMBEREDLIST"],
		TokenTypes["LIST"],
		TokenTypes["WORD"],
		TokenTypes["SPACE"],
		TokenTypes["SPECIALCHAR"],
	}
	if token := P.Match(RecursiveTypes); token != EmptyToken {
		n := P.ParseText()
		return Node{operator: token, operand: []*Node{&n}}
	}
	if token := P.Match([]TokenType{TokenTypes["LINE"]}); token != EmptyToken {
		return Node{operator: token, operand: nil}
	}
	if token := P.Match([]TokenType{TokenTypes["CODE"]}); token != EmptyToken {
		CodeNodes := []*Node{}
		for P.Match([]TokenType{TokenTypes["CODE"]}) == EmptyToken {
			n := P.ParseText()
			CodeNodes = append(CodeNodes, &n)
		}
		return Node{operator: token, operand: CodeNodes}
	}
	var node = Node{operator: Token{}, operand: []*Node{}}
	for P.Match([]TokenType{TokenTypes["SEMICOLON"]}) == EmptyToken {
		n := P.ParseText()
		node.operand = append(node.operand, &n)
	}
	if token := P.Match([]TokenType{TokenTypes["SEMICOLON"]}); token != EmptyToken {
		return Node{operator: token, operand: []*Node{}}
	}
	return node
}

func (P *Parser) ParseList() Node {
	listnode := Node{operator: Token{Type: SecondTokenTypes["GROUPNUMBEREDLIST"], Text: "GROUPNUMBEREDLIST"}}
	token := P.Match([]TokenType{TokenTypes["NUMBEREDLIST"]})
	for token != EmptyToken {
		node := Node{operator: token, operand: []*Node{}}
		for P.Match([]TokenType{TokenTypes["SEMICOLON"]}) != EmptyToken {
			n := P.ParseText()
			node.operand = append(node.operand, &n)
		}
		listnode.operand = append(listnode.operand, &node)
	}
	return listnode
}

func (P *Parser) ParseText() Node {
	RecursiveTypes := []TokenType{
		TokenTypes["SPECIALCHAR"],
		TokenTypes["WORD"],
		TokenTypes["SPACE"],
		TokenTypes["ITALIC"],
		TokenTypes["BOLT"],
	}
	if token := P.Match(RecursiveTypes); token != EmptyToken {
		n := P.ParseText()
		return Node{operator: token, operand: []*Node{&n}}
	}
	return Node{}
}
