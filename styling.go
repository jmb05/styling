package styling

import "fmt"

const RESET = "\033[0m"

const (
	black = iota
	red
	green
	yellow
	blue
	purple
	cyan
	white
	defaultStyle
	bold
	dim
	italic
	underline
	blinking
	reverse
	background
	highIntensity
	boldHighIntensity
	backgroundHighIntensity
)

func Style(color uint8, style uint8, s string) string {
	styleInt, colorAdd := styleStringToInt(style)
	return create(styleInt, colorStringToInt(color), colorAdd) + s + RESET
}

func create(style uint8, color uint8, colorAdd uint8) string {
	return fmt.Sprintf("\033[%d;%dm", style, color+colorAdd)
}

func styleStringToInt(style uint8) (uint8, uint8) {
	switch style {
	case defaultStyle, bold, dim, italic, underline, blinking, reverse:
		return style - 8, 0
	case background:
		return 0, 10
	case highIntensity:
		return 0, 60
	case boldHighIntensity:
		return 1, 60
	case backgroundHighIntensity:
		return 0, 70
	default:
		return 0, 0
	}
}

func colorStringToInt(color uint8) uint8 {
	return color + 30
}
