package utils

import "strings"
func ToSmallCaps(s string) string {
	// Map of regular letters to their small caps Unicode equivalents
	smallCaps := map[rune]rune{
		'a': 'ᴀ', 'b': 'ʙ', 'c': 'ᴄ', 'd': 'ᴅ', 'e': 'ᴇ',
		'f': 'ғ', 'g': 'ɢ', 'h': 'ʜ', 'i': 'ɪ', 'j': 'ᴊ',
		'k': 'ᴋ', 'l': 'ʟ', 'm': 'ᴍ', 'n': 'ɴ', 'o': 'ᴏ',
		'p': 'ᴘ', 'q': 'ǫ', 'r': 'ʀ', 's': 's', 't': 'ᴛ',
		'u': 'ᴜ', 'v': 'ᴠ', 'w': 'ᴡ', 'x': 'x', 'y': 'ʏ',
		'z': 'ᴢ',
	}

	var result []rune
	for _, r := range s {
		if val, ok := smallCaps[r]; ok {
			result = append(result, val)
		} else if val, ok := smallCaps[lower(r)]; ok {
			result = append(result, val)
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func lower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}



func ToCapsLock(s string) string {
	return strings.ToUpper(s)
}


func TruncateWithEllipsis(s string, limit int) string {
	// Replace '|' with '-'
	s = strings.ReplaceAll(s, "|", "-")

	// Truncate with ellipsis if needed
	if len(s) <= limit {
		return s
	}
	if limit <= 3 {
		return s[:limit] // Not enough room for "..."
	}
	return s[:limit-3] + "..."
}