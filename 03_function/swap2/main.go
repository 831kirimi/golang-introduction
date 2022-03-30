package main

func swap2(n, m *int) {
	*n, *m = *m, *n
}
func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}
