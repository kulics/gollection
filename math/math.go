package math

type Integer interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 |
		~uint64 | ~uint32 | ~uint16 | ~uint8
}

type Number interface {
	Integer | ~float64 | ~float32
}

func Inc[T Number](a T) T {
	return a + 1
}

func Dec[T Number](a T) T {
	return a - 1
}

func Add[T Number](a T) func(b T) T {
	return func(b T) T {
		return a + b
	}
}

func ToAdd[T Number](a T) func(b T) T {
	return func(b T) T {
		return b + a
	}
}

func Sub[T Number](a T) func(b T) T {
	return func(b T) T {
		return a - b
	}
}

func ToSub[T Number](a T) func(b T) T {
	return func(b T) T {
		return b - a
	}
}

func Mul[T Number](a T) func(b T) T {
	return func(b T) T {
		return a * b
	}
}

func ToMul[T Number](a T) func(b T) T {
	return func(b T) T {
		return b * a
	}
}

func Div[T Number](a T) func(b T) T {
	return func(b T) T {
		return a / b
	}
}

func ToDiv[T Number](a T) func(b T) T {
	return func(b T) T {
		return b / a
	}
}
