package formatters

// FormatPunctuation formats punctuation with proper spacing
// Punctuation sticks to previous word, space after
// Note: Spacing is handled by FSM, this just returns the punctuation
func FormatPunctuation(punct string) string {
	// For now, just return as-is since FSM handles spacing
	// Punctuation groups (e.g., "...") are already combined by FSM
	return punct
}
