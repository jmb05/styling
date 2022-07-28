# styling
A ANSI terminal styling library for go

## Installation
Run `go get github.com/jmb05/styling` in your go project folder

## Usage
Examples:
```
package main

import ( 
  "github.com/jmb05/styling"
  "fmt"
)

func main() {
  fmt.Println(styling.Color("red", "Some Red Text"))
  fmt.Println(styling.ColorBold("green", "Some BOLD green text"))
  fmt.Println(styling.ColorDim("vlue", "Some dimmed purple text"))
  fmt.Println(styling.ColorBackground("yellow", "Some text on a yellow background"))
}

``` 
Available styles `Color ColorBold ColorDim ColorItalic ColorUnderline ColorBlinking ColorReverse ColorBackground ColorHighIntensity ColorBoldHighIntensity ColorBackgroundHighIntensity`

NOTE: Some styles might not work in your terminal emulator

Available colors `black red green yellow blue purple cyan white`
