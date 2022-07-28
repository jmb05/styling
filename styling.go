package styling

import "fmt"

const RESET = "\033[0m"

func Style(color string, style string, s string) string {
	styleInt, colorAdd := styleStringToInt(style)
	return create(styleInt, colorStringToInt(color), colorAdd) + s + RESET
}

func create(style uint8, color uint8, colorAdd uint8) string {
	return fmt.Sprintf("\033[%d;%dm", style, color+colorAdd)
}

func styleStringToInt(style string) (uint8, uint8) {
	switch style {
	case "bold":
		return 1, 0
	case "dim":
		return 2, 0
	case "italic":
		return 3, 0
	case "underline":
		return 4, 0
	case "blinking":
		return 5, 0
	case "reverse":
		return 7, 0
	case "background":
		return 0, 10
	case "high_intensity":
		return 0, 60
	case "bold_high_intensity":
		return 1, 60
	case "background_high_intensity":
		return 0, 70
	default:
		return 0, 0
	}
}

func colorStringToInt(color string) uint8 {
	switch color {
	case "black":
		return 30
	case "red":
		return 31
	case "green":
		return 32
	case "yellow":
		return 33
	case "blue":
		return 34
	case "purple":
		return 35
	case "cyan":
		return 36
	default:
		return 37
	}
}
