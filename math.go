package main

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
