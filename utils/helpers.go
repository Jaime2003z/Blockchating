package utils

// ValidateLength verifica que una cadena no exceda un límite de longitud
func ValidateLength(input string, maxLength int) bool {
	return len(input) <= maxLength
}
