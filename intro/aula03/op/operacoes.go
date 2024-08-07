package op

// Soma função começando com maisculo é exportada ..
func Soma(a int, b int) int {
	return a + b
}

// sub não exportada, começando por minusculo ...
func sub(a int) int {
	return a - 1
}

// Sub funçãndo será exportada ...
func Sub(a int) int {
	return sub(a)
}
