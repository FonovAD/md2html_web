package MarkdownToHTML

const (
	HTMLsizeMultiplier = 5
	HTMLsizeDevisor    = 4
	HTMLPrefix         = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body style="margin-left: 3vw; margin-top: 2vh;">`
	HTMLPostfix = `
	</body>
	</html>`
)
