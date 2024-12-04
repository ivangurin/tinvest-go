package utils

type Set[T comparable] map[T]struct{}

// ToSet преобразует слайс collection в map[key]struct{} по полю
func ToSet[TKey comparable, TArg any](collection []TArg, selector func(arg TArg) TKey) Set[TKey] {
	result := make(map[TKey]struct{}, len(collection))
	for _, item := range collection {
		result[selector(item)] = struct{}{}
	}
	return result
}

// Add добавляет указанный элемент в множество.
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// Has проверяет есть ли указанный элемент.
func (s Set[T]) Has(elem T) bool {
	_, ok := s[elem]

	return ok
}

// Delete удаляем элемент
func (s Set[T]) Delete(elem T) {
	delete(s, elem)
}

// ToSlice возвращает слайс из элементов множества.
func (s Set[T]) ToSlice() []T {
	if s == nil {
		return nil
	}

	ret := make([]T, 0, len(s))
	for val := range s {
		ret = append(ret, val)
	}

	return ret
}
