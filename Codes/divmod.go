package piscine

func DivMod(a int, b int, div *int, mod *int) {
	result := a / b
	remainder := a % b
	*div = *div + result
	*mod = *mod + remainder
}
