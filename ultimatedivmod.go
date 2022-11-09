package piscine

func UltimateDivMod(a *int, b *int) {
	ent := *a / (*b)
	rst := *a % (*b)
	*a = ent
	*b = rst
}
