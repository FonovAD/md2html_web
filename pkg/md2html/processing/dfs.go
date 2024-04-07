package processing

import "bytes"

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
			HTMLLine.WriteString("<hr stule=\"border: none; background-color: black; color: black; height: 2px;\"></hr>")
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
			HTMLLine.WriteString(`<li style="list-style-type:'`)
			HTMLLine.WriteString(node.operator.Text)
			HTMLLine.WriteString(`'; margin-left:1vw">`)
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
			HTMLLine.WriteString("</li>")
		}
	case "CODE":
		{
			HTMLLine.WriteString("<code>")
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
			HTMLLine.WriteString("</code>")
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
			HTMLLine.WriteString(`<li style="margin-left:1vw">`)
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
			HTMLLine.WriteString("</li>")
		}
	case "ITALIC":
		{
			HTMLLine.WriteString("<i>")
			HTMLLine.WriteString(node.operator.Text[1 : len(node.operator.Text)-1])
			HTMLLine.WriteString("</i>")
			for i := range node.operand {
				HTMLLine.WriteString(LineLayout(*node.operand[i]))
			}
		}
	case "BOLT":
		{
			HTMLLine.WriteString("<b>")
			HTMLLine.WriteString(node.operator.Text[2 : len(node.operator.Text)-2])
			HTMLLine.WriteString("</b>")
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
