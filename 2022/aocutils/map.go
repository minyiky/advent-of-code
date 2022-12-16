package aocutils

func CopyMap[A comparable, B any](a map[A]B) map[A]B {
	b := make(map[A]B)
	for k, v := range a {
		b[k] = v
	}
	return b
}
