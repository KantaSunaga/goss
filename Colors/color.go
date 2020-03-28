package Colors


type ColorText struct {
	r	 int
	g 	 int
	b 	 int
	Text string
}

func Red(text string) ColorText{
	red := ColorText{ r: 255, g: 0, b: 0, Text: text }
	return red
}

func Green(text string) ColorText{
	green := ColorText{ r: 51, g: 255, b: 0, Text: text }
	return green
}

func Blue(text string) ColorText{
	blue := ColorText{ r: 0, g: 0, b: 139, Text: text }
	return blue
}