package yqlib

import (
	"fmt"
)

// Thanks @risentveber!

const escape = "\x1b"

// Attribute defines a single SGR Code
type Attribute int

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

func format(attr Attribute) string {
	return fmt.Sprintf("%s[%dm", escape, attr)
}
