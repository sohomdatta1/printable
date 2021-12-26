// Go library to check if a byte is a printable character.
package printable

import "strings"

// All printable characters
var AllPrintable string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&\\'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"

// Layout changing characters
var LayoutChanging string = "\t\n\r\x0b\x0c"

// Check if a byte is a printable character.
func IsPrintable(b byte) bool {
	return strings.IndexByte(AllPrintable, b) != -1
}

// Check if a byte is a printable character, excluding layout changing characters.
func IsPrintableAndLayoutSafe(b byte) bool {
	return strings.IndexByte(AllPrintable, b) != -1 && strings.IndexByte(LayoutChanging, b) == -1
}
