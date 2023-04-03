package generics

func SumInts(m map[string]int) int {
	var s int
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K comparable, V int | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

//func Sum[K comparable, V addable](m map[K]V) V {
//	var s V
//	for _, v := range m {
//		s += v
//	}
//	return s
//}
//
//type addable interface {
//	constraints.Number
//	+ ~int | ~float64
//}
