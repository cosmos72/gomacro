package main

func toInt(x interface{}) (int64, bool) {
	switch x := x.(type) {
	case int64:
		return x, true
	case uint64:
		return int64(x), true
	default:
		return 0, false
	}
}

func toUint(x interface{}) (uint64, bool) {
	switch x := x.(type) {
	case int64:
		return uint64(x), true
	case uint64:
		return x, true
	default:
		return 0, false
	}
}

func toFloat(x interface{}) (float64, bool) {
	switch x := x.(type) {
	case int64:
		return float64(x), true
	case uint64:
		return float64(x), true
	case float64:
		return x, true
	default:
		return 0.0, false
	}
}

func toComplex(x interface{}) (complex128, bool) {
	switch x := x.(type) {
	case complex128:
		return x, true
	default:
		f, ok := toFloat(x)
		return complex(f, 0.0), ok
	}
}
