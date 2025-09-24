package yqlib

import (
	"fmt"

	"github.com/fatih/color"
)

// Thanks @risentveber!

const escape = "\x1b"

func format(attr color.Attribute) string {
	return fmt.Sprintf("%s[%dm", escape, attr)
}
