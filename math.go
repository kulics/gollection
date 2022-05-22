package gollection

// Provide functional increase
func Inc[T Number](a T) T {
	return a + 1
}

// Provide functional decrease
func Dec[T Number](a T) T {
	return a - 1
}

// Provide functional add
func Add[T Number](a T) func(b T) T {
	return func(b T) T {
		return a + b
	}
}

// Provide functional add, the first parameter is the right operand
func ToAdd[T Number](a T) func(b T) T {
	return func(b T) T {
		return b + a
	}
}

// Provide functional sub
func Sub[T Number](a T) func(b T) T {
	return func(b T) T {
		return a - b
	}
}

// Provide functional sub, the first parameter is the right operand
func ToSub[T Number](a T) func(b T) T {
	return func(b T) T {
		return b - a
	}
}

// Provide functional mul
func Mul[T Number](a T) func(b T) T {
	return func(b T) T {
		return a * b
	}
}

// Provide functional mul, the first parameter is the right operand
func ToMul[T Number](a T) func(b T) T {
	return func(b T) T {
		return b * a
	}
}

// Provide functional div
func Div[T Number](a T) func(b T) T {
	return func(b T) T {
		return a / b
	}
}

// Provide functional div, the first parameter is the right operand
func ToDiv[T Number](a T) func(b T) T {
	return func(b T) T {
		return b / a
	}
}
