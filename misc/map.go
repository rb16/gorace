package main

func AddMap() {
	m := map[int]int{0: 1, 1: 1}
	_ = m[0] + m[1]
}

func AddStruct() {
	m := struct {
		a int
		b int
	}{0, 1}
	_ = m.a + m.b
}
func main() {
	AddMap()
	AddStruct()
}
