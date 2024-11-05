package styling

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	fmt.Println(Style(white, bold, "White Bold"))
	fmt.Println(Style(red, highIntensity, "Red High Intensity"))
	fmt.Println(Style(green, boldHighIntensity, "Green Bold High Intensity"))
}
