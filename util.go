package main

func ContainsValue[K comparable, V comparable](m map[K]V, target V) bool {
	for _, value := range m {
		if value == target {
			return true
		}
	}
	return false
}

func CheckValue[K comparable, V comparable, T any](m map[K]V, target T, checker func(V, T) bool) bool {
	for _, value := range m {
		if checker(value, target) {
			return true
		}
	}
	return false
}

func FindKeyByValue[K comparable, V comparable](m map[K]V, target V) (K, bool) {
	var zeroKey K
	for key, value := range m {
		if value == target {
			return key, true
		}
	}
	return zeroKey, false
}

func FindAllKeysByValue[K comparable, V comparable](m map[K]V, target V) []K {
	var keys []K
	for key, value := range m {
		if value == target {
			keys = append(keys, key)
		}
	}
	return keys
}
