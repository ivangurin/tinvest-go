package utils

func Filter[T any](collection []T, selector func(arg T) bool) []T {
	res := make([]T, 0, len(collection))
	for _, item := range collection {
		if selector(item) {
			res = append(res, item)
		}
	}
	return res
}

// Distinct Возвращает уникальные значения из массива
func Distinct[T comparable](s []T) []T {
	if len(s) < 2 {
		return Clone(s)
	}

	dict := make(map[T]struct{})
	for _, element := range s {
		dict[element] = struct{}{}
	}
	res := make([]T, 0, len(dict))
	for key := range dict {
		res = append(res, key)
	}
	return res
}

// Contains reports whether v is present in s.
func Contains[T comparable](slice []T, value T) bool {
	return Index(slice, value) >= 0
}

// ToMap преобразует slice в map
func ToMap[TKey comparable, TValue, TArg any](collection []TArg, selector func(arg TArg) (TKey, TValue)) map[TKey]TValue {
	result := make(map[TKey]TValue, len(collection))
	for _, element := range collection {
		k, v := selector(element)
		result[k] = v
	}
	return result
}

// ToMapByField индексирует слайс в map по полю
func ToMapByField[TKey comparable, TArg any](collection []TArg, selector func(arg TArg) TKey) map[TKey]TArg {
	result := make(map[TKey]TArg, len(collection))
	for _, item := range collection {
		result[selector(item)] = item
	}
	return result
}

// ChunkBy разбивает слайс на чанки
func ChunkBy[T any](items []T, chunkSize int) [][]T {
	chunks := make([][]T, 0, len(items)/chunkSize+1)

	if len(items) == 0 {
		return chunks
	}

	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[S ~[]E, E any](s S) S {
	// Preserve nil in case it matters.
	if s == nil {
		return nil
	}
	return append(S([]E{}), s...)
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func Grow[S ~[]E, E any](s S, n int) S {
	if n < 0 {
		panic("cannot be negative")
	}
	if n -= cap(s) - len(s); n > 0 {
		return append(s, make(S, n)...)[:len(s)]
	}
	return s
}

// Select создает слайс из полей объектов исходного слайса
func Select[Tr, TArg any](collection []TArg, f func(arg TArg) Tr) []Tr {
	result := make([]Tr, len(collection))
	for i, element := range collection {
		result[i] = f(element)
	}
	return result
}

// GroupBy - сгруппировать объекты по ключу
func GroupBy[TKey comparable, TValue any](collection []TValue, f func(arg TValue) TKey) map[TKey][]TValue {
	result := make(map[TKey][]TValue, len(collection))
	for _, element := range collection {
		k := f(element)
		result[k] = append(result[k], element)
	}
	return result
}

// Skip panic-safe src[n:]
func Skip[T any](src []T, n int) []T {
	if n >= len(src) {
		return []T{}
	}
	return src[n:]
}

// Take panic-safe src[:n]
func Take[T any](src []T, n int) []T {
	if n >= len(src) {
		return src
	}
	return src[:n]
}
