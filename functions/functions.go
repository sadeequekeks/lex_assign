package functions

func Addit(x *int, y *int) int {
	return *x + *y
}

func Subtrac(x *int, y *int) int {
	return *x - *y
}

func Multi(x *int, y *int) int {
	return *x * *y
}

func Divis(x *float64, y *float64) float64 {
	return *x / *y
}
