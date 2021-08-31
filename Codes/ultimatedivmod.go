package piscine

func UltimateDivMod(a *int, b *int) {
	result := *a / *b
	remaider := *a % *b
	*a = result
	*b = remaider
}
