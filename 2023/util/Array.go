package util

func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T any, U any](array []T, reducer func(acc U, val T) U, initialValue U) U {
	result := initialValue
	for _, value := range array {
		result = reducer(result, value)
	}
	return result
}