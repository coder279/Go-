package main

import "log"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	log.Println(r)
	return r
}
func main() {
	m := make(map[string]interface{})
	m["name"] = "Leo"
	m["age"] = 1
	MapKeys(m)

}
