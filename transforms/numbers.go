package transforms

import (
	"strconv"
	"strings"
)

// HexToDec converts hexadecimal string to decimal string
func HexToDec(hex string) string {
	hex = strings.TrimSpace(hex)
	if hex == "" {
		return ""
	}

	val, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return hex // Return original if invalid
	}
	return strconv.FormatInt(val, 10)
}

// BinToDec converts binary string to decimal string
func BinToDec(bin string) string {
	bin = strings.TrimSpace(bin)
	if bin == "" {
		return ""
	}

	val, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return bin // Return original if invalid
	}
	return strconv.FormatInt(val, 10)
}
