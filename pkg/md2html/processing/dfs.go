package processing

import "bytes"

const (
	LINEprefix          = "<hr stule=\"border: none; background-color: black; color: black; height: 2px;\"></hr>"
	NUMBEREDLISTprefix1 = `<li style="list-style-type:'`
	NUMBEREDLISTprefix2 = `'; margin-left:1vw">`
	NUMBEREDLISTpostfix = "</li>"
	CODEprefix          = "<code>"
	CODEpostfix         = "</code>"
	LISTprefix          = `<li style="margin-left:1vw">`
	LISTpostfix         = "</li>"
	ITALICprefix        = "<i>"
	ITALICpostfix       = "</i>"
	BOLTprefix          = "<b>"
	BOLTpostfix         = "</b>"
)

func Run(node StatmentsNode, file_size int) string {
	var HTML bytes.Buffer
	HTML.Grow(file_size)
	for i := range node.CodeString {
		HTML.WriteString(LineLayout(node.CodeString[i]))
	}
	return HTML.String()
}

func LineLayout(node Node) string {
	var HTMLLine bytes.Buffer
	switch node.operator.Type.name {
	case "HEADING":
		{
			HTMLLine.WriteString(prefixHeadings[node.operator.Text])
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
			HTMLLine.WriteString(postfixHeadings[node.operator.Text])
		}
	case "LINE":
		{
			HTMLLine.WriteString(LINEprefix)
		}
	case "SEMICOLON": // It will never happen
		{ // Reserved for future feature additions
			HTMLLine.WriteString("\n")
		}
	case "WORD":
		{
			HTMLLine.WriteString(node.operator.Text)
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
		}
	case "NUMBEREDLIST":
		{
			HTMLLine.WriteString(NUMBEREDLISTprefix1)
			HTMLLine.WriteString(node.operator.Text)
			HTMLLine.WriteString(NUMBEREDLISTprefix2)
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
			HTMLLine.WriteString(NUMBEREDLISTpostfix)
		}
	case "CODE":
		{
			HTMLLine.WriteString(CODEprefix)
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
			HTMLLine.WriteString(CODEpostfix)
		}
	case "SPACE":
		{
			HTMLLine.WriteString(" ")
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
		}
	case "LIST":
		{
			HTMLLine.WriteString(LISTprefix)
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
			HTMLLine.WriteString(LISTpostfix)
		}
	case "ITALIC":
		{
			HTMLLine.WriteString(ITALICprefix)
			HTMLLine.WriteString(node.operator.Text[1 : len(node.operator.Text)-1])
			HTMLLine.WriteString(ITALICpostfix)
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
		}
	case "BOLT":
		{
			HTMLLine.WriteString(BOLTprefix)
			HTMLLine.WriteString(node.operator.Text[2 : len(node.operator.Text)-2])
			HTMLLine.WriteString(BOLTpostfix)
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
		}
	case "SPECIALCHAR":
		{
			HTMLLine.WriteString(node.operator.Text)
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
		}
	}
	return HTMLLine.String()
}
