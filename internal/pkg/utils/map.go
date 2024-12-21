package utils

func ToMap[TKey comparable, TValue, TArg any](collection []TArg, selector func(arg TArg) (TKey, TValue)) map[TKey]TValue {
	result := make(map[TKey]TValue, len(collection))
	for _, element := range collection {
		k, v := selector(element)
		result[k] = v
	}
	return result
}
